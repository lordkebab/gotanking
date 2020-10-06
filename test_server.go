package gotanking

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	client *WOTClient
	server *httptest.Server
)

// ServerSetup instantiates a new test
func ServerSetup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = NewClient("dummy", SetBaseURL(server.URL))

	return func() {
		server.Close()
	}
}

// Fixture returns the test data for the path endpoint
func Fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}

	return string(b)
}
