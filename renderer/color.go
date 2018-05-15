package renderer

import "image/color"

type ColorFunc func(int, int, int, int, int) color.Color

var Colors = map[string]ColorFunc{
	"black": Black,
	"reds":  Reds,
	"white": White,
}

func (fn ColorFunc) Color(x, xOffset, y, yOffset, height int) color.Color {
	return fn(x, xOffset, y, yOffset, height)
}

func (fn ColorFunc) Palette() color.Palette {
	p := color.Palette{color.White, color.Black}
	for i := 0; i < 255; i += 10 {
		p = append(p, color.RGBA{R: uint8(255 - i), G: 0, B: 0, A: 255})
	}
	return p
}

func White(x, xOffset, y, yOffset, height int) color.Color {
	return color.White
}

func Black(x, xOffset, y, yOffset, height int) color.Color {
	return color.Black
}

func Reds(x, xOffset, y, yOffset, height int) color.Color {
	return color.RGBA{R: uint8(255 - y*255/height), G: 0, B: 0, A: 255}
}

func RegColorFunc(name string, fn ColorFunc) {
	Colors[name] = fn
}
