package main

import (
	"encoding/json"
	"net/http"

	"github.com/wwgberlin/go-funk/audio"
)

func samplesHandler(data audio.WavData) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.Encode(data.Samples)
	}
}
