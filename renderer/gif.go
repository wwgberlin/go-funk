package renderer

import (
	"image"
	"image/gif"
	"time"
)

func MakeGif(points []int, width, height int, duration time.Duration, colorFunc ColorFunc) *gif.GIF {
	delay := int(duration.Seconds() / float64(len(points)) * 1000 / 10) // delay between frames (10ms)

	anim := gif.GIF{LoopCount: len(points)}

	points = project(points, height)
	for _, v := range points {
		img := renderFrame(v, width, height, colorFunc)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	return &anim
}

func renderFrame(v int, width, height int, colorFunc ColorFunc) *image.Paletted {
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, colorFunc.Palette())

	DrawRectangle(img, 0, v, width, height, colorFunc)

	return img
}
