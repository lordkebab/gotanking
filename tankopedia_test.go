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

	arena, err := client.ListMaps(nil)
	if err != nil {
		t.Error(err)
	}

	if len(arena.Data) == 0 {
		t.Errorf("No maps returned")
	}
}

func TestDisplayFilter(t *testing.T) {
	testServer := ServerSetup()
	defer testServer()

	client, _ := NewClient("dummy", SetBaseURL(server.URL))

	mux.HandleFunc("/encyclopedia/arenas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, Fixture("encyclopedia/arena_display_filter.json"))

	})

	displayOpts := &MapInput{
		Fields: []string{
			"camouflage_type",
		},
		Language: "en",
	}
	arena, err := client.ListMaps(displayOpts)
	if err != nil {
		t.Error(err)
	}

	got := arena.Data["04_himmelsdorf"].Camo
	want := "summer"

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
