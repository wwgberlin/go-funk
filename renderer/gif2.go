package renderer

import (
	"image"
	"image/gif"
	"time"
)

func MakeGif2(points []int, width, height int, duration time.Duration, colorFunc ColorFunc) *gif.GIF {
	delay := int(duration.Seconds() / float64(len(points)) * 1000 / 10) // delay between frames (10ms)

	anim := gif.GIF{LoopCount: len(points)}

	points = project(points, height)

	for i := range points {
		low, high := clampMin(i-width/2, 0), clampMax(i+width/2, len(points))
		img := renderFrame2(points[low:high], width, height, colorFunc)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	return &anim

}

func renderFrame2(points []int, width, height int, colorFunc ColorFunc) *image.Paletted {
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, colorFunc.Palette())

	for x, y := range points {
		DrawRectangle(img, x, y, 1, height, colorFunc)
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
