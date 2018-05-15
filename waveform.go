package main

import (
	"fmt"
	"image/png"
	"math"
	"net/http"
	"strconv"

	"github.com/wwgberlin/go-funk/audio"
	"github.com/wwgberlin/go-funk/renderer"
	"github.com/wwgberlin/go-funk/sampler"
)

func waveformHandler(data audio.WavData) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		width, err := strconv.Atoi(r.URL.Query().Get("width"))
		if err != nil {
			width = 1200
		}

		samplingKey := r.URL.Query().Get("sampling")

		samplingFunc, ok := sampler.Samplers[samplingKey]
		if !ok {
			http.Error(w, fmt.Sprintf("could not find sampling method with key %q", samplingKey), http.StatusInternalServerError)
			return
		}

		colorKey := r.URL.Query().Get("colors")
		colorFunc, ok := renderer.Colors[colorKey]
		if !ok {
			http.Error(w, fmt.Sprintf("could not find colors method with key %q", colorKey), http.StatusInternalServerError)
			return
		}

		compressed, err := sampler.Compress(data.Samples, width, samplingFunc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		rgba := renderer.DrawWaveform(compressed, math.MaxUint8, colorFunc)
		if err = png.Encode(w, rgba); err != nil {
			panic(err)
		}

	}
}
