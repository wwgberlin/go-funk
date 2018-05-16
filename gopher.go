package main

import (
	"image"
	"image/color"

	"github.com/wwgberlin/go-funk/renderer"
)

// ColorGopherFunc takes an image and returns a renderer.ColorFunc
// that takes x,xOffset, y, yOffset and height, calculates the original
// x and y (respective to the new height) and returns the color at the original
// co-ordinates.
// Use the ratio between the original image height and the given height
// to determine the original x and y, then call img.At(origX, origY) to
// return the color.

// Example:
// given imgHeight = 100 and height = 50
// the ratio between the original image and the new image is then 2
// if x = 15, xOffset = 10 the original x would be 10 (15-10)*2

func ColorGopherFunc(img image.Image) renderer.ColorFunc {
	// to get the original image height uncomment this:
	//imgHeight := img.Bounds().Max.Y - img.Bounds().Min.Y
	return func(x, xOffset, y, yOffset, height int) color.Color {
		return color.Black
	}
}

var ColorGopherPalette = color.Palette{
	color.White,
	color.Black,
	color.RGBA{156, 202, 217, 255},
	color.RGBA{255, 215, 54, 255},
	color.RGBA{125, 125, 125, 255},
}
