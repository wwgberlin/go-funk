package renderer

import (
	"image"
	"math"
)

func DrawWaveform(points []int, height int, colorFunc ColorFunc) *image.RGBA {
	points = project(points, height)
	width, height := len(points), max(points)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	DrawRectangle(img, 0, height, width, height, White)

	for x, y := range points {
		DrawRectangle(img, x, y, 1, height, colorFunc)
	}
	return img
}

func max(points []int) int {
	m := math.MinInt64
	for _, p := range points {
		if p > m {
			m = p
		}
	}
	return m
}

func normalize(v, height, maxSlice int) int {
	return v * height / maxSlice
}

//todo: make public add tests and remove implementation
func project(in []int, height int) []int {
	maxSlice := max(in)
	out := make([]int, len(in))
	for i, v := range in {
		out[i] = normalize(v, height, maxSlice)
	}
	return out
}
