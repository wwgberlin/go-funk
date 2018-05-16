package renderer

import (
	"image"
	"image/gif"
	"io"
	"time"
)

func MakeGif2(w io.Writer, points []int, duration time.Duration, conf Config) {
	delay := int(duration.Seconds() / float64(len(points)) * 1000 / 10) // delay between frames (10ms)

	anim := gif.GIF{LoopCount: len(points)}

	for i := range points {
		low, high := clampMin(i-conf.Width/2, 0), clampMax(i+conf.Width/2, len(points))
		img := renderFrame2(points[low:high], conf.Width, conf.Height, conf.Colorer)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(w, &anim)
}

func renderFrame2(points []int, width, height int, colorer Colorer) *image.Paletted {
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, colorer.Palette())

	for x, y := range points {
		DrawRectangle(img, x, height-y, x+1, height, height, colorer.Fn)
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
