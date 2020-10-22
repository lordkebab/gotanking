package gotanking

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAccount(t *testing.T) {
	testServer := ServerSetup()
	defer testServer()

	client, _ := NewClient("dummy", SetBaseURL(server.URL))

	mux.HandleFunc("/account/list/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, Fixture("account/list.json"))
	})

	search := "lulz_man"
	input := &AccountInput{
		Language: "en",
	}

	_, err := client.GetAccount(search, input)

	if err != nil {
		t.Error(err)
	}
}

func TestGetAccountID(t *testing.T) {
	testServer := ServerSetup()
	defer testServer()

	client, _ := NewClient("dummy", SetBaseURL(server.URL))
	mux.HandleFunc("/account/list/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, Fixture("account/list.json"))
	})

	search := "lulz_man"
	accountNumber := client.GetAccountID(search)
	want := 1008273454

	if accountNumber != want {
		t.Errorf("got %d want %d", accountNumber, want)
	}
}
