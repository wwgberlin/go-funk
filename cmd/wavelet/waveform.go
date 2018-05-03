package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/soundcloud/wavelet/lib/sampler"
	"github.com/soundcloud/wavelet/lib/wav"
)

func waveform(w http.ResponseWriter, req *http.Request) {
	width, err := strconv.Atoi(req.URL.Query().Get("width"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	points, err := sampler.Sample(samples, width)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	drawWaveform(w, points)
}

func drawWaveform(w io.Writer, points []int) {
	width, height := len(points), 255
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	setBackground(img, color.White)

	for i, v := range points {
		barHeight := clamp(v, 255)
		for j := height; j > height-barHeight; j-- {
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

func clamp(v, max int) int {
	return int(math.Min(math.Abs(float64(v)), float64(max)))
}
