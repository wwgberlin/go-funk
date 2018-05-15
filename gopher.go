package main

import (
	"image"
	"image/color"
	"image/gif"
	"net/http"
	"strconv"

	"github.com/wwgberlin/go-funk/audio"
	"github.com/wwgberlin/go-funk/renderer"
	"github.com/wwgberlin/go-funk/sampler"
)

func gopherHandler(data audio.WavData, img image.Image) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		width, err := strconv.Atoi(req.URL.Query().Get("width"))
		if err != nil {
			width = 100
		}
		height, err := strconv.Atoi(req.URL.Query().Get("height"))
		if err != nil {
			height = 100
		}
		count, err := strconv.Atoi(req.URL.Query().Get("count"))
		if err != nil {
			count = 3000
		}

		compressed, err := sampler.Compress(data.Samples, count, sampler.AbsAvg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		anim := renderer.Gopher(compressed, width, height, data.Duration, colorGopherFunc(img))

		gif.EncodeAll(w, anim)
	}
}

func colorGopherFunc(img image.Image) renderer.ColorFunc {
	return func(x int, y int, height int) color.Color {
		ratio := float64(img.Bounds().Max.Y) / float64(height)
		oldX := ratio * float64(x)
		oldY := ratio * float64(y)
		c := img.At(int(oldX), int(oldY))
		return c
	}
}
