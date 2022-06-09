package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// main result: map[_id:ObjectID("620291e8a55b7866b72f547f") appearances:41 average_discount:10.317073170731707 average_price:22.253414634146353
type DbProduct struct {
	Title            string
	Appearances      int32
	Average_discount float64
	Average_price    float64
}

type PenguinProduct struct {
	Title              string
	Description        string
	OriginalPrice      float64
	DiscountPrice      float64
	DiscountPercentage float64
	Rating             int64
	IsValid            bool
	Reason             string
}

// Connection URI
var uri string

// getProduct will get the current product from penguin magic
func getProduct() []byte {
	// make http request to get product
	res, err := http.Get(os.Getenv("API_URL") + "/coinProduct")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)

	return body
}

// constructProductObj construct a product object from the mongodb result
func constructProductObj(b bson.M) DbProduct {
	product := DbProduct{Title: b["title"].(string), Appearances: b["appearances"].(int32), Average_discount: b["average_discount"].(float64), Average_price: b["average_price"].(float64)}
	return product
}

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
	fmt.Println("Successfully connected and pinged.")
	return client
}

func main() {
	client := connectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// var p Product
	var penguinProduct PenguinProduct
	p := getProduct()
	if err := json.Unmarshal(p, &penguinProduct); err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("main penguinProduct: %v", penguinProduct)) // __AUTO_GENERATED_PRINT_VAR__

	coll := client.Database("penguin_magic").Collection("open_box")
	var result bson.M
	if err := coll.FindOne(context.TODO(), bson.D{{"title", penguinProduct.Title}}).Decode(&result); err != nil {
		panic(err)
	}

	dbProduct := constructProductObj(result)
	fmt.Println(fmt.Sprintf("main dbProduct: %+v", dbProduct)) // __AUTO_GENERATED_PRINT_VAR__
	updateProduct(&dbProduct, penguinProduct)
	fmt.Println(fmt.Sprintf("main dbProduct: %v", dbProduct.Appearances)) // __AUTO_GENERATED_PRINT_VAR__

}

// updateProduct update the dbproduct with the new product passed in, and return an up to date product
func updateProduct(dbProduct *DbProduct, penguinProduct PenguinProduct) error {
	if dbProduct.Title != penguinProduct.Title {
		// return an error
		return fmt.Errorf("Product titles do not match")
	}
	// update the dbproduct with the new product
	dbProduct.Appearances = dbProduct.Appearances + 1
	dbProduct.Average_discount = (dbProduct.Average_discount*float64(dbProduct.Appearances-1) + penguinProduct.DiscountPrice) / float64(dbProduct.Appearances)
	dbProduct.Average_price = (dbProduct.Average_price*float64(dbProduct.Appearances-1) + penguinProduct.DiscountPrice) / float64(dbProduct.Appearances)

	return nil
}
