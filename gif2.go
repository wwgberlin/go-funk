package main

import (
	"fmt"
	"image/gif"
	"net/http"
	"strconv"

	"github.com/wwgberlin/go-funk/audio"
	"github.com/wwgberlin/go-funk/renderer"
	"github.com/wwgberlin/go-funk/sampler"
)

func gif2Handler(data audio.WavData) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		width, err := strconv.Atoi(req.URL.Query().Get("width"))
		if err != nil {
			width = 50
		}
		height, err := strconv.Atoi(req.URL.Query().Get("height"))
		if err != nil {
			height = 100
		}
		count, err := strconv.Atoi(req.URL.Query().Get("count"))
		if err != nil {
			count = 3000
		}

		colorKey := req.URL.Query().Get("colors")
		colorFunc, ok := renderer.Colors[colorKey]
		if !ok {
			http.Error(w, fmt.Sprintf("could not find colors method with key %q", colorKey), http.StatusInternalServerError)
			return
		}

		compressed, err := sampler.Compress(data.Samples, count, sampler.AbsAvg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		anim := renderer.MakeGif2(compressed, width, height, data.Duration, colorFunc)
		gif.EncodeAll(w, anim)
	}
}
