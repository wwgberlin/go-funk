package renderer

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"time"
)

func Gopher(w io.Writer, points []int, duration time.Duration, conf Config) {
	delay := int(duration.Seconds() / float64(len(points)) * 1000 / 10) // delay between frames (10ms)

	anim := gif.GIF{LoopCount: len(points)}

	points = Project(points, conf.Height)

	for _, v := range points {
		img := renderGopher(v, conf.Width, conf.Height, conf.ColorFunc)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(w, &anim)
}

func renderGopher(v int, width, height int, colorFunc ColorFunc) *image.Paletted {
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, color.Palette{
		color.White,
		color.Black,
		color.RGBA{156, 202, 217, 255},
		color.RGBA{255, 215, 54, 255},
		color.RGBA{125, 125, 125, 255},
	})

	x1 := (width - v) / 2
	y1 := (height - v) / 2

	DrawRectangle(img, 0, 0, width, height, height, White)
	DrawRectangle(img, (width-v)/2, (height-v)/2, x1+v, y1+v, v, colorFunc)

	return img
}
