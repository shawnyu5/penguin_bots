package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Product struct {
	Title              string
	Description        string
	OriginalPrice      float64
	DiscountPrice      float64
	DiscountPercentage float64
	Rating             int64
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

func main() {
	godotenv.Load()
	uri = os.Getenv("MONGODB_URI")
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	var productStruct Product
	product := getProduct()
	if err := json.Unmarshal(product, &productStruct); err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("main product: %v", productStruct)) // __AUTO_GENERATED_PRINT_VAR__
	// coll := client.Database("penguin_magic").Collection("open_box")
	// var result bson.M
	// err = coll.FindOne(context.TODO(), bson.D{{"title", "Self Tying Shoelace by Jay Noblezada"}}).Decode(&result)

	// fmt.Println(fmt.Sprintf("main result: %v", result)) // __AUTO_GENERATED_PRINT_VAR__

}
