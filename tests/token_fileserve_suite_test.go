package token_fileserve

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"net/http"
	"github.com/fanszhub/token-fileserve/fileServers"
)

func TestTokenFileserve(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TokenFileserve Suite")
}

var _ = Describe("Token", func() {
	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				var handler http.Handler = token_fileserve.NewTokenFileServer([]string{"string", "string2"},"A dir")
				Expect(handler).ToNot(Equal(nil))
			})
		})
	})
})