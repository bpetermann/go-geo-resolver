package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	Port = ":4000"
)

var (
	citiesDir = "data"
)

type Cities struct {
	Cities []City `json:"cities"`
}

type City struct {
	Name string  `json:"city"`
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
}

func resolveCity(city string) (*City, error) {
	rootPath, _ := os.Getwd()

	cityFilePath := filepath.Join(rootPath, citiesDir, fmt.Sprintf("%s.json", strings.ToLower(city[0:1])))

	jsonFile, err := os.Open(cityFilePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var cities Cities
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&cities)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(cities.Cities); i++ {
		if strings.EqualFold(cities.Cities[i].Name, city) {
			return &cities.Cities[i], nil
		}
	}

	return nil, fmt.Errorf("City not found")
}

func geocode(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Mehod Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City Not Found", http.StatusNotFound)
		return
	}

	cityData, err := resolveCity(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cityData)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/geocode/", geocode)

	log.Print("Server running on", Port)
	err := http.ListenAndServe(Port, mux)
	log.Fatal(err)
}
