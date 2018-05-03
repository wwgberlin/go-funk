package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/wwgberlin/go-funk/lib/audio"
	"github.com/wwgberlin/go-funk/lib/sampler"
)

func points(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("fixtures/fixture.wav")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	samples, err := audio.Samples(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	points, err := sampler.Avg(samples, 1200)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(points)
}
