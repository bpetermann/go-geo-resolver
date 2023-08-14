package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Cities struct {
	Cities []City `json:"cities"`
}

type City struct {
	Name string  `json:"city"`
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
}

func resolveCity(city string) ([]byte, error) {
	jsonFile, err := os.Open("cities.json")

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var cities Cities

	err = json.Unmarshal(byteValue, &cities)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(cities.Cities); i++ {
		if strings.ToLower(cities.Cities[i].Name) == city {
			result, err := json.Marshal(cities.Cities[i])
			if err != nil {
				return nil, err
			}
			return result, nil
		}
	}

	return nil, nil
}

func home(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Mehod Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	city := r.URL.Query().Get("city")

	cityJson, err := resolveCity(strings.ToLower(city))

	if err == nil && cityJson != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(cityJson)
	} else {
		http.Error(w, "City Not Found", http.StatusNotFound)
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Server running on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
