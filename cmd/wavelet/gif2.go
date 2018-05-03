package main

import (
	"image"
	"image/gif"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/wwgberlin/go-funk/lib/audio"
	"github.com/wwgberlin/go-funk/lib/sampler"
)

func gifHandler2(w http.ResponseWriter, req *http.Request) {
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

	drawWaveformGif2(w, points, width, height)
}

func drawWaveformGif2(w io.Writer, points []int, width, height int) {
	const (
		duration = 213.574649 // track duration in seconds
	)

	delay := int(duration / float64(len(points)) * 1000 / 10) // delay between frames (10ms)

	anim := gif.GIF{LoopCount: len(points)}

	for i := range points {
		low, high := clampMin(i-width/2, 0), clampMax(i+width/2, len(points))
		img := renderFrame2(points[low:high], width, height)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(w, &anim)
}

func renderFrame2(points []int, width, height int) *image.Paletted {
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)

	for x, v := range points {
		for y := height; y > height-normalize(v, height, 255); y-- {
			img.SetColorIndex(x, y, blackIndex)
		}
	}

	return img
}

func clampMin(v, min int) int {
	if v < min {
		return min
	}
	return v
}

func clampMax(v, max int) int {
	if v > max {
		return max
	}
	return v
}
