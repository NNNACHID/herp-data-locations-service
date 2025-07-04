package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Observation struct {
	ID   int    `json:"id"`
	Name string `json:"species_guess"`
	Lat  float64 `json:"latitude"`
	Lon  float64 `json:"longitude"`
	Date string `json:"observed_on"`
}

type APIResponse struct {
	Results []Observation `json:"results"`
}

func main() {
	url := "https://api.inaturalist.org/v1/observations?taxon_name=Python&per_page=5"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erreur GET: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data APIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Erreur JSON: %v", err)
	}

	for _, obs := range data.Results {
		fmt.Printf("ID: %d | Nom: %s | Lat: %.4f | Lon: %.4f | Date: %s\n", obs.ID, obs.Name, obs.Lat, obs.Lon, obs.Date)
	}
}
