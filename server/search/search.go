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
	Original_price      float64 `json:"original_price"`
	Discount_price      float64 `json:"discount_price"`
	Discount_percentage float64 `json:"discount_percentage"`
	Appearances         int32   `json:"appearances"`
}

// func (product *Product) fillDefaults() {
// product.Title = ""
// product.Appearances = 0
// }

func main() {
	client := connectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		log.Println("Database successfully disconnected.")
	}()

	// coll := client.Database("penguin_magic").Collection("open_box")
	SearchByRegex(&Product{Title: "card"})
}

// connectDB will connect to the mongodb
func connectDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
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
// Returns an bson array of products
func SearchByRegex(product *Product) []bson.D {
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

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	// for _, result := range results {
	// fmt.Println(result)
	// }
	return results
}
