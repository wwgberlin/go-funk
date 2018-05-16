package renderer

import (
	"bytes"
	"fmt"
	"image/color"
	"image/gif"
	"testing"
	"time"
)

func TestRenderGopherFrame(t *testing.T) {
	img := RenderGopherFrame(0, 3, 3, Colors["black"])
	if img == nil {
		t.Fatalf("RenderGopherFrame was expected to return a *image.Paletted")
	}

	for i := 0; i < 3; i++ {
		if img.At(i, i) != color.White {
			t.Fatalf("Unexpected color at (%d,%d) wanted white", i, i)
		}
	}
	img = RenderGopherFrame(1, 3, 3, Colors["black"])

	for i := 0; i < 3; i++ {
		if i == 1 {
			if img.At(i, i) != color.Black {
				t.Fatalf("Unexpected color at (%d,%d) wanted white", i, i)
			}
		} else {
			if img.At(i, i) != color.White {
				t.Fatalf("Unexpected color at (%d,%d) wanted white", i, i)
			}
		}
	}
}
func TestGopher(t *testing.T) {
	var b bytes.Buffer

	h := 3
	Gopher(&b, []int{0, 0, 1, 1}, time.Second, Config{Colorer: Colors["black"], Width: h, Height: h, Count: 4})

	img, err := gif.DecodeAll(&b)
	if err != nil {
		t.Fatal("Decoding the retreived image was expected to succeed")
	}

	l := 4
	if len(img.Delay) != l {
		t.Fatalf("number of delayes was expected to be %d but got: %d", l, len(img.Delay))
	}
	d := 25
	for i, v := range img.Delay {
		if v != d {
			t.Fatalf("delay %d returned %d but expected %d", i, v, d)
		}
	}

	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	//First 2 frames - v = 0 - all white
	for _, v := range img.Image[0:2] {
		for i := 0; i < h; i++ {
			if v.At(i, i) != white {
				t.Fatalf("Unexpected color at (%d,%d) wanted white", i, i)
			}
		}
	}

	//Last 2 frames - v = 0 - 1 pixel in the middle must be black
	for _, v := range img.Image[2:4] {
		for i := 0; i < h; i++ {
			if i == 1 {
				if v.At(i, i) != black {
					fmt.Println(i, i, v.At(i, i))
					t.Fatalf("Unexpected color at (%d,%d) wanted white", i, i)
				}
			} else {
				if v.At(i, i) != white {
					fmt.Println(i, i, v.At(i, i))
					t.Fatalf("Unexpected color at (%d,%d) wanted white", i, i)
				}
			}
		}
	}

}
