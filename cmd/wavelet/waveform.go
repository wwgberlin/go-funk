package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/soundcloud/wavelet/lib/sampler"
	"github.com/soundcloud/wavelet/lib/utils"
	"github.com/soundcloud/wavelet/lib/wav"
)

var methods = map[string]sampler.SamplerFunc{
	"avg":     sampler.AvgSample,
	"abs_avg": sampler.AbsAvgSample,
}

func waveform(w http.ResponseWriter, req *http.Request) {
	width, err := strconv.Atoi(req.URL.Query().Get("width"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	methodKey := req.URL.Query().Get("method")
	if methodKey == "" {
		methodKey = "abs_avg"
	}

	samplingMethod, ok := methods[methodKey]
	if !ok {
		http.Error(w, fmt.Sprintf("could not find sampling method with key %q", methodKey), http.StatusInternalServerError)
		return
	}

	file, err := os.Open("fixtures/fixture.wav")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	samples, err := wav.Samples(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	points, err := samplingMethod(samples, width)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	drawWaveform(w, points)
}

func drawWaveform(w io.Writer, points []int) {
	s := utils.IntSliceStats(points)
	width, height := len(points), s.Max
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	setBackground(img, color.White)

	for i, v := range points {
		for j := height; j > height-v; j-- {
			img.Set(i, j, color.Black)
		}
	}

	png.Encode(w, img)
}

func setBackground(img *image.RGBA, c color.Color) {
	width, height := img.Rect.Size().X, img.Rect.Size().Y
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, c)
		}
	}
}
