package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
	print()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Server running on :4000")
	err := http.ListenAndServe(":400", mux)
	log.Fatal(err)
}
