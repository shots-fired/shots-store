package store_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shots-fired/shots-store/store"
)

func TestStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Store Suite")
}

var _ = Describe("Store", func() {
	It("New store is not nil", func() {
		Expect(store.New()).ToNot(BeNil())
	})
})
