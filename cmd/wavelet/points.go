package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/soundcloud/wavelet/lib/audio"
	"github.com/soundcloud/wavelet/lib/sampler"
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
	defer file.Close()

	points, err := sampler.Avg(samples, 1200)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(points)
}
