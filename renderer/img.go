package renderer

import (
	"image/color"
)

type Img interface {
	Set(x, y int, c color.Color)
}

// Draw Rectangle receives and image and draws a rectangle
// iterate over x and y axis:
// x: from xOffset to width
// y: from height to height-yOffset
// Use colorFunc(x,y,height) to determine the color to paint each pixel
// To color a pixel use img.Set

func DrawRectangle(img Img, xOffset, yOffset, width, height int, colorFunc ColorFunc) {
	for x := 0; x < width; x++ {
		for y := height; y > height-yOffset; y-- {
			c := colorFunc(x, y, height)
			// todo: check that we are not exceeding the image here and
			// fix the tests
			// the core library tends to swallow this kind of error
			img.Set(x+xOffset, y, c)
		}
	}
}
