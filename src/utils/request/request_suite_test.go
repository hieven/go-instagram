package request

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ts        *httptest.Server
	tsHandler func(http.ResponseWriter, *http.Request)
)

func TestRequest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Request Suite")
}

var _ = BeforeSuite(func() {
	tsHandler = func(http.ResponseWriter, *http.Request) {}

	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tsHandler(w, r)
	}))
})

var _ = AfterSuite(func() {
	ts.Close()
})
