package main

import (
	"log"
	"net/http"
)

type City struct {
	name string
	lat  float64
	lng  float64
}

func home(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Mehod Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	city := r.URL.Query().Get("city")

	if city == "paris" {
		var parisJson = []byte(`{"name": "Paris","lat": 48.8566, "lng":  2.3522 }`)
		w.Header().Set("Content-Type", "application/json")
		w.Write(parisJson)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Server running on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
