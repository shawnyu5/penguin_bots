package search

import (
	"context"
	"testing"
)

var ss = SearchServiceImpl{}

// TestAbleToConnectToDb tests if we can connect to the database
func TestAbleToConnectToDb(t *testing.T) {
	client := ss.connectDB()
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
	found, err := ss.SearchByRegex(product)
	if err != nil {
		t.Error("Error searching for product")
	}

	if len(found) == 0 {
		t.Error("Failed to find product")
	}
}
