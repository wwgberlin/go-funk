package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/wwgberlin/go-funk/lib/audio"
	"github.com/wwgberlin/go-funk/lib/sampler"
)

func points(w http.ResponseWriter, req *http.Request) {
	count, err := strconv.Atoi(req.URL.Query().Get("count"))
	if err != nil {
		count = 2000
	}

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

	points, err := sampler.Avg(samples, count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(points)
}
