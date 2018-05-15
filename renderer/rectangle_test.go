package renderer

import (
	"image/color"
	"testing"
)

type pos struct {
	x, y int
}

type imgMock map[pos]color.Color

func (img *imgMock) Set(x, y int, c color.Color) {
	(*img)[pos{x, y}] = c
}

func mockColorFunc() ColorFunc {
	prev := color.White
	return func(i int, i2 int, i3, i4, i5 int) color.Color {
		if prev == color.White {
			prev = color.Black
		} else {
			prev = color.White
		}
		return prev
	}
}

func TestDrawRectangle(t *testing.T) {
	var m imgMock = make(map[pos]color.Color)
	expected := imgMock{
		pos{10, 10}: color.Black,
		pos{10, 11}: color.White,
		pos{11, 10}: color.Black,
		pos{11, 11}: color.White,
	}

	DrawRectangle(&m, 10, 10, 12, 12, 12, mockColorFunc())

	if len(m) != len(expected) {
		t.Fatalf("Unexpected rectangle number of pixles: %d. expected: %d", len(m), len(expected))
	}
	for k, v := range m {
		if v2, ok := expected[k]; !ok {
			t.Fatalf("Position %v was not expected to be written to", k)
		} else if v2 != v {
			t.Fatalf("Unexpected color %v in position %v. Expected: %v", v, k, v2)
		}
	}
}
