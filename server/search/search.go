package search

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Product struct {
	Title               string  `json:"title"`
	Average_discount    float64 `json:"average_discount"`
	Average_price       float64 `json:"average_price"`
	Discount_percentage float64 `json:"discount_percentage"`
	Appearances         int32   `json:"appearances"`
}

// func (product *Product) fillDefaults() {
// product.Title = ""
// product.Appearances = 0
// }

// func main() {
// client := connectDB()
// defer func() {
// if err := client.Disconnect(context.TODO()); err != nil {
// panic(err)
// }
// log.Println("Database successfully disconnected.")
// }()

// // coll := client.Database("penguin_magic").Collection("open_box")
// SearchByRegex(&Product{Title: "card"})
// }

// connectDB will connect to the mongodb
func connectDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	fmt.Println(fmt.Sprintf("connectDB uri: %v", uri)) // __AUTO_GENERATED_PRINT_VAR__
	if uri == "" {
		panic("MONGODB_URI is not set")
	}
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

// SearchByRegex will search for a product title by a regex
// Returns an array of products
func SearchByRegex(product *Product) []Product {
	client := connectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		log.Println("Database successfully disconnected.")
	}()

	coll := client.Database("penguin_magic").Collection("open_box")
	filter := bson.M{}

	filter["title"] = bson.M{"$regex": fmt.Sprintf(".*%s.*", product.Title)}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var products []Product
	// go through all products and add to array of products
	for cursor.Next(context.TODO()) {
		var singleProduct Product
		err := cursor.Decode(&singleProduct)
		if err != nil {
			panic(err)
		}
		products = append(products, singleProduct)
	}

	// for _, result := range results {
	// fmt.Println(result)
	// current := result[0]
	// products = append(products, Product{Title: string(current)})
	// }
	return products
}
