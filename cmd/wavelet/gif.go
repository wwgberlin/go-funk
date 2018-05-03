package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/wwgberlin/go-funk/lib/audio"
	"github.com/wwgberlin/go-funk/lib/sampler"
)

func gifHandler(w http.ResponseWriter, req *http.Request) {
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

	file, err := os.Open("fixtures/fixture.wav")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	samples, err := audio.Samples(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	points, err := sampler.AbsAvg(samples, count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	drawWaveformGif(w, points, width, height)
}

func drawWaveformGif(w io.Writer, points []int, width, height int) {
	const (
		duration = 213.574649 // track duration in seconds
	)

	delay := int(duration / float64(len(points)) * 1000 / 10) // delay between frames (10ms)

	anim := gif.GIF{LoopCount: len(points)}

	for _, v := range points {
		img := renderFrame(v, width, height)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(w, &anim)
}

var palette = []color.Color{color.Black, color.White}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func renderFrame(v int, width, height int) *image.Paletted {
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)

	for x := 0; x < width; x++ {
		for y := height; y > height-normalize(v, height, 255); y-- {
			img.SetColorIndex(x, y, blackIndex)
		}
	}

	return img
}

func normalize(v, height, max int) int {
	return v * height / max
}
