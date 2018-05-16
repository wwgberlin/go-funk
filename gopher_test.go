package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"testing"
	"time"
)

type XY struct {
	x, y int
}

type mockImage struct {
	image.Image
	dim int
	at  func(x, y int) color.Color
}

func (m mockImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.dim, m.dim)
}

func (m mockImage) At(x, y int) color.Color {
	return m.at(x, y)
}

func TestColorGopherFunc(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	dim := rand.Intn(100) + 1
	newDim := rand.Intn(dim) + 1

	recX, recY := math.MinInt64, math.MinInt64
	res := color.Gray{Y: uint8(rand.Intn(255))}

	m := mockImage{dim: dim, at: func(x, y int) color.Color {
		recX, recY = x, y
		return res
	}}

	x, y := rand.Intn(newDim), rand.Intn(newDim)
	xOffset, yOffset := 1, 1
	c := ColorGopherFunc(m).Color(x, xOffset, y, yOffset, newDim)

	if c != res {
		t.Fatalf("ColorGopherFunc returned Unexpected color. Exepected: %v. Got: %v", res, c)
	}

	expectedX := (x - xOffset) * dim / newDim
	expectedY := (y - yOffset) * dim / newDim

	if recX != expectedX || recY != expectedY {
		t.Fatalf("Unexpected request to At. Expected At(%d,%d) but got At(%d,%d)",
			expectedX, expectedY, recX, recY)
	}
}
