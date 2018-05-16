package renderer

import (
	"image"
	"image/png"
	"io"
	"time"
)

func DrawWaveform(w io.Writer, points []int, _ time.Duration, conf Config) {
	width, height := len(points), max(points)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	DrawRectangle(img, 0, 0, width, height, height, White)

	for x, y := range points {
		DrawRectangle(img, x, height-y, x+1, height, height, conf.Colorer.Color)
	}
	png.Encode(w, img)
}
