package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Logger", func() {
	AppFs = afero.NewMemMapFs()
	Context("CacheProduct", func() {
		It("Should cache a product", func() {
			title := "hello world"
			product := PenguinProduct{
				Title: title,
			}
			CacheProduct(product)

			f, err := afero.ReadFile(AppFs, "product.txt")
			Expect(err).To(BeNil())
			Expect(string(f)).To(BeEquivalentTo(title))
		})
	})

})
