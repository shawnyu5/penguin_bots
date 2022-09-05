package search

import (
	"context"
	"log"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Search", func() {
	var ss SearchService
	BeforeEach(func() {
		logger := log.New(os.Stdout, "", log.LUTC)
		ss = SearchServiceImpl{}
		ss = LoggingMiddleware{Logger: logger, Next: ss}
	})

	It("Should be able to connect to the database", func() {
		client := ss.connectDB()
		defer func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()

		Expect(client).NotTo(BeNil(), "Failed to connect to mongodb")
	})

	Context("SearchByRegex", func() {
		It("Should be able to search for a product", func() {
			product := &Product{Title: "card"}
			found, err := ss.SearchByRegex(product)
			Expect(err).To(BeNil(), "Error searching for product")
			Expect(len(found)).NotTo(Equal(0), "Failed to find product")
		})
	})
})
