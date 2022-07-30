package search

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoggingMiddleware struct {
	Logger *log.Logger
	Next   SearchService
}

func (lm LoggingMiddleware) SearchByRegex(p *Product) ([]Product, error) {
	lm.Logger.Printf("(SearchByRegex) Searching for product: %s", p.Title)
	return lm.Next.SearchByRegex(p)
}

func (lm LoggingMiddleware) connectDB() *mongo.Client {
	lm.Logger.Println("Connecting to database")
	defer func() {
		lm.Logger.Println("Disconnecting from database")
	}()
	return lm.Next.connectDB()
}
