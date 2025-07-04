package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	observations := main()
	if err != nil {
		t.Fatalf("Erreur API: %v", err)
	}

	if len(observations) == 0 {
		t.Fatalf("Aucune observation reçue")
	}

	for _, obs := range observations {
		t.Logf("ID: %d | Nom: %s | Lat: %.4f | Lon: %.4f | Date: %s", obs.ID, obs.Name, obs.Lat, obs.Lon, obs.Date)
	}
}
