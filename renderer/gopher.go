package renderer

import (
	"image"
	"image/gif"
	"io"
	"time"
)

// Gopher writes a new gif into w with delays of trackDuration normalized to
// the length so that the delay between frames is the duration of the
// track divided by the length of points by 10ms
// for each element in points, call RenderGopherFrame
// append the returned frame (image) into gif.Image
// and append the the delay for the frame

func Gopher(w io.Writer, points []int, duration time.Duration, conf Config) {
	delay := int(duration.Seconds() / float64(len(points)) * 1000 / 10) // delay between frames (10ms)

	anim := gif.GIF{LoopCount: len(points)}

	for _, v := range points {
		img := RenderGopherFrame(v, conf.Width, conf.Height, conf.Colorer)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(w, &anim)
}

// RenderGopherFrame returns a new paletted image of rectangle(0,0,width, height)
// and uses DrawRectangle to draw a square in the middle of the image
// of height=width=v using colorer.
// Example:
// given v = 10 width=100, height=120 the following params will be passed to
// DrawRectangle:
// x1: 100-10/2 = 45, x2: x1+v=55
// y1: 120-10/2	= 55, y2: y1+v=65
// * hint: look at examples in the code to see usages of
// image.NewPaletted and image.Rect(...)

func RenderGopherFrame(v int, width, height int, colorer Colorer) *image.Paletted {
	return nil
}
