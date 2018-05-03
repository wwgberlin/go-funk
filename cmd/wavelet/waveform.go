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

	"github.com/soundcloud/wavelet/lib/audio"
	"github.com/soundcloud/wavelet/lib/sampler"
	"github.com/soundcloud/wavelet/lib/utils"
)

var sampling = map[string]sampler.SamplerFunc{
	"avg":     sampler.Avg,
	"abs_avg": sampler.AbsAvg,
}

type ColorFunc func(int, int) color.Color

var colors = map[string]ColorFunc{
	"simple": simple,
	"reds":   reds,
}

func waveform(w http.ResponseWriter, req *http.Request) {
	width, err := strconv.Atoi(req.URL.Query().Get("width"))
	if err != nil {
		width = 1200
	}

	samplingKey := req.URL.Query().Get("sampling")
	if samplingKey == "" {
		samplingKey = "abs_avg"
	}
	samplingFunc, ok := sampling[samplingKey]
	if !ok {
		http.Error(w, fmt.Sprintf("could not find sampling method with key %q", samplingKey), http.StatusInternalServerError)
		return
	}

	colorKey := req.URL.Query().Get("colors")
	if colorKey == "" {
		colorKey = "simple"
	}
	colorFunc, ok := colors[colorKey]
	if !ok {
		http.Error(w, fmt.Sprintf("could not find colors method with key %q", colorKey), http.StatusInternalServerError)
		return
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

	points, err := samplingFunc(samples, width)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	drawWaveform(w, points, colorFunc)
}

func drawWaveform(w io.Writer, points []int, colorFunc ColorFunc) {
	s := utils.IntSliceStats(points)
	width, height := len(points), s.Max
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	setBackground(img, color.White)

	for i, v := range points {
		for j := height; j > height-v; j-- {
			img.Set(i, j, colorFunc(i, j))
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

func simple(x, y int) color.Color {
	return color.Black
}

func reds(x, y int) color.Color {
	return color.RGBA{R: uint8(255 - y), G: 0, B: 0, A: 255}
}
