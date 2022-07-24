package search

import (
	"context"
	"testing"
)

// TestAbleToConnectToDb tests if we can connect to the database
func TestAbleToConnectToDb(t *testing.T) {
	client := connectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if client == nil {
		t.Error("Failed to connect to mongodb")
	}
}

func TestSearchByRegex(t *testing.T) {
	product := &Product{Title: "card"}
	found := SearchByRegex(product)
	if len(found) == 0 {
		t.Error("Failed to find product")
	}
}
