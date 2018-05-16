package renderer

import (
	"image"
	"image/gif"
	"io"
	"time"
)

func Gopher(w io.Writer, points []int, duration time.Duration, conf Config) {
	delay := int(duration.Seconds() / float64(len(points)) * 1000 / 10) // delay between frames (10ms)

	anim := gif.GIF{LoopCount: len(points)}

	for _, v := range points {
		img := renderGopher(v, conf.Width, conf.Height, conf.Colorer)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(w, &anim)
}

func renderGopher(v int, width, height int, colorer Colorer) *image.Paletted {
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, colorer.Palette())

	x1 := (width - v) / 2
	y1 := (height - v) / 2

	DrawRectangle(img, 0, 0, width, height, height, White)
	DrawRectangle(img, (width-v)/2, (height-v)/2, x1+v, y1+v, v, colorer.Color)

	return img
}
