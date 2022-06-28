package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DbProduct struct {
	Title            string  `json:"title"`
	Appearances      int32   `json:"appearances"`
	Average_discount float64 `json:"average_discount"`
	Average_price    float64 `json:"average_price"`
	Created_at       string  `json:"created_date"`
	Updated_at       string  `json:"updated_date"`
}

type PenguinProduct struct {
	Title               string  `json:"title"`
	Original_price      float64 `json:"original_price"`
	Discount_price      float64 `json:"discount_price"`
	Discount_percentage float64 `json:"discount_percentage"`
}

// Connection URI
var uri string
var c *cache.Cache

func main() {
	c = cache.New(cache.NoExpiration, 10*time.Minute)
	client := connectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	for {
		Log(client)
		time.Sleep(2 * time.Minute)
	}
}

// Log initiate the logger
func Log(client *mongo.Client) {

	penguinProduct := getProductInfo()
	hasChanged := hasProductChanged(penguinProduct)
	// if product has not changed, do nothing
	// We want to log when the product has changed. the last time it was seen
	if !hasChanged {
		log.Printf("Product %s has not changed", penguinProduct.Title)
		return
	}

	cacheProduct(penguinProduct)

	coll := client.Database("penguin_magic").Collection("open_box")

	var dbResult bson.M
	// find the product with current product title in db
	err := coll.FindOne(context.TODO(), bson.D{{"title", penguinProduct.Title}}).Decode(&dbResult)
	// create default product if product is not found
	if err == mongo.ErrNoDocuments {
		const appearances int32 = 0
		const average_discount float64 = 0
		const average_price float64 = 0
		var created_at string = time.Now().Format("2006-01-02 15:04:05")
		var updated_at string = time.Now().Format("2006-01-02 15:04:05")

		dbResult = bson.M{"title": penguinProduct.Title, "appearances": appearances, "average_discount": average_discount, "average_price": average_price, "created_at": created_at, "updated_at": updated_at}
	}
	dbProduct := constructProductObj(dbResult)

	err = updateProduct(&dbProduct, penguinProduct)
	if err != nil {
		panic(err)
	}
	saveProduct(&dbProduct, coll)
}

// hasProductChanged will check if the current product title is different from the cached product title.
// Returns true if the product has changed.
func hasProductChanged(product PenguinProduct) bool {
	cacheProduct, found := c.Get("product_title")
	if !found {
		// if not found, assume product has not changed to avoid counting too many times
		return false
	}

	// if product is same, product has not changed
	if cacheProduct.(string) == product.Title {
		return false
	}
	return true
}

// cacheProduct cache product title under "product_title" key
func cacheProduct(product PenguinProduct) {
	c.Set("product_title", product.Title, cache.NoExpiration)
}

// getProductInfo will get the current product from penguin magic. Returns a penguin product object
func getProductInfo() PenguinProduct {
	// make http request to get product
	res, err := http.Get(os.Getenv("API_URL") + "/logger")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)

	penguinProduct := PenguinProduct{}
	if err := json.Unmarshal(body, &penguinProduct); err != nil {
		panic(err)
	}

	return penguinProduct
}

// constructProductObj construct a product object from the mongodb result
func constructProductObj(b bson.M) DbProduct {
	product := DbProduct{Title: b["title"].(string), Appearances: b["appearances"].(int32), Average_discount: b["average_discount"].(float64), Average_price: b["average_price"].(float64)}

	// not all products has these time stamps
	if b["created_at"] == nil {
		product.Created_at = time.Now().Format("2006-01-02 15:04:05")
	} else {
		product.Created_at = b["created_at"].(string)
	}

	if b["updated_at"] == nil {
		product.Updated_at = time.Now().Format("2006-01-02 15:04:05")
	} else {
		product.Updated_at = b["updated_at"].(string)

	}
	return product
}

// connectDB will connect to the mongodb
func connectDB() *mongo.Client {
	godotenv.Load()
	uri = os.Getenv("MONGODB_URI")
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("Database successfully connected and pinged.")
	return client
}

// saveProduct save the dbproduct to the collection
func saveProduct(dbProduct *DbProduct, coll *mongo.Collection) {
	b, err := bson.Marshal(dbProduct)
	if err != nil {
		panic(err)
	}

	filter := bson.D{{"title", dbProduct.Title}}
	result, err := coll.ReplaceOne(context.TODO(), filter, b)
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("saveProduct result: %v", result.UpsertedID)) // __AUTO_GENERATED_PRINT_VAR__

}

// updateProduct update the dbproduct with the new product passed in
func updateProduct(dbProduct *DbProduct, penguinProduct PenguinProduct) error {
	if dbProduct.Title != penguinProduct.Title {
		// return an error
		return fmt.Errorf("Product titles do not match")
	}

	// update the dbproduct with the new product
	dbProduct.Appearances = dbProduct.Appearances + 1

	dbProduct.Average_discount = (dbProduct.Average_discount*float64(dbProduct.Appearances-1) + penguinProduct.Discount_percentage) / float64(dbProduct.Appearances)

	dbProduct.Average_price = (dbProduct.Average_price*float64(dbProduct.Appearances-1) + penguinProduct.Discount_price) / float64(dbProduct.Appearances)
	// get current date and time
	dbProduct.Updated_at = time.Now().Format("2006-01-02 15:04:05")

	return nil
}
