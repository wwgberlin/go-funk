package main

import (
	"encoding/json"
	"image"
	"image/color"
	"math"
	"net/http"

	"github.com/wwgberlin/go-funk/audio"
	"github.com/wwgberlin/go-funk/renderer"
	"github.com/wwgberlin/go-funk/sampler"
)

func handler(data audio.WavData, defaults *RequestData, rdr renderer.RenderFunc) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := parseRequest(r, defaults)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		compressed, err := sampler.Compress(data.Samples, reqData.Conf.Count, reqData.Sampler)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rdr(w, compressed, data.Duration, reqData.Conf)
	}
}

func samplesHandler(data audio.WavData) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.Encode(data.Samples)
	}
}

func gifHandler(data audio.WavData) func(w http.ResponseWriter, req *http.Request) {
	defaults := RequestData{
		Sampler: sampler.AbsAvg,
		Conf: renderer.Config{
			Width:     50,
			Height:    100,
			Count:     3000,
			ColorFunc: renderer.Black,
		},
	}
	return handler(data, &defaults, renderer.MakeGif)
}

func gif2Handler(data audio.WavData) func(w http.ResponseWriter, req *http.Request) {

	defaults := RequestData{
		Sampler: sampler.AbsAvg,
		Conf: renderer.Config{
			Width:     50,
			Height:    100,
			Count:     3000,
			ColorFunc: renderer.Black,
		},
	}
	return handler(data, &defaults, renderer.MakeGif2)
}

func gopherHandler(data audio.WavData, img image.Image) func(w http.ResponseWriter, req *http.Request) {
	defaults := RequestData{
		Sampler: sampler.AbsAvg,
		Conf: renderer.Config{
			Width:     100,
			Height:    100,
			Count:     3000,
			ColorFunc: ColorGopherFunc(img),
		},
	}
	return handler(data, &defaults, renderer.Gopher)
}

func waveformHandler(data audio.WavData) func(w http.ResponseWriter, r *http.Request) {
	defaults := RequestData{
		Sampler: sampler.AbsAvg,
		Conf: renderer.Config{
			Width:     1200,
			Count:     1200,
			Height:    math.MaxUint8,
			ColorFunc: renderer.Black,
		},
	}

	return handler(data, &defaults, renderer.DrawWaveform)
}

// ColorGopherFunc takes an image and returns a renderer.ColorFunc
// that takes x,y and height, calculates the original x and y
// (respective to the new height) and returns the color at the original
// co-ordinates.
// Use the ratio between the original image height and the given height
// to determine the original x and y, then call img.At(origX, origY) to
// return the color.

func ColorGopherFunc(img image.Image) renderer.ColorFunc {
	imgHeight := img.Bounds().Max.Y - img.Bounds().Min.Y
	return func(x, xOffset, y, yOffset, height int) color.Color {
		oldX := (x - xOffset) * imgHeight / height
		oldY := (y - yOffset) * imgHeight / height
		return img.At(oldX, oldY)
	}
}
