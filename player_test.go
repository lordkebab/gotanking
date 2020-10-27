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

func TestGetPlayerPersonalData(t *testing.T) {
	testServer := ServerSetup()
	defer testServer()

	client, _ := NewClient("dummy", SetBaseURL(server.URL))
	mux.HandleFunc("/account/info/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, Fixture("account/info.json"))
	})

	resp, _ := client.GetPlayerPersonalData(123, nil)

	want := 8165
	treesCut := resp.Data["1008273454"].Statistics.TreesCut

	if treesCut != want {
		t.Errorf("Got %q want %q", treesCut, want)
	}
}

func TestGetPlayerVehicles(t *testing.T) {
	testServer := ServerSetup()
	defer testServer()

	client, _ := NewClient("dummy", SetBaseURL(server.URL))
	mux.HandleFunc("/account/tanks/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, Fixture("account/tanks.json"))
	})

	accountID := 123

	_, err := client.GetPlayerVehicles(accountID, nil)
	if err != nil {
		t.Error(err)
	}
}
