package renderer

import (
	"image/color"
	"testing"
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
				t.Fatalf("Unexpected color at (%d,%d) wanted black", i, i)
			}
		} else {
			if img.At(i, i) != color.White {
				t.Fatalf("Unexpected color at (%d,%d) wanted white", i, i)
			}
		}
	}
}
