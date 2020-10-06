package gotanking

import (
	"fmt"
	"net/http"
	"testing"
)

func TestListMaps(t *testing.T) {
	testServer := ServerSetup()
	defer testServer()

	client, _ := NewClient("dummy", SetBaseURL(server.URL))

	mux.HandleFunc("/encyclopedia/arenas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, Fixture("encyclopedia/arenas.json"))

	})

	_, err := client.ListMaps()
	if err != nil {
		t.Errorf("Error: %q", err.Error())
	}
}
