package renderer

import "image/color"

type (
	ColorFunc func(x, xOffset, y, yOffset, height int) color.Color

	Colorer struct {
		Fn ColorFunc
		P  color.Palette
	}
)

var Colors = map[string]Colorer{
	"black": {Black, color.Palette{color.White, color.Black}},
	"white": {White, color.Palette{color.White, color.Black}},
	"reds":  {Reds, redsPalette()},
}

func (fn ColorFunc) Color(x, xOffset, y, yOffset, height int) color.Color {
	return fn(x, xOffset, y, yOffset, height)
}
func (c Colorer) Color(x, xOffset, y, yOffset, height int) color.Color {
	return c.Fn.Color(x, xOffset, y, yOffset, height)
}

func (c Colorer) Palette() color.Palette {
	return c.P
}

func redsPalette() color.Palette {
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

func RegColorFunc(name string, colorer Colorer) {
	Colors[name] = colorer
}
