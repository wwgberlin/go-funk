package wav

import (
	"math"
	"path/filepath"
	"testing"
)

func TestSamples(t *testing.T) {
	file, err := filepath.Abs("../../fixtures/fixture.wav")
	if err != nil {
		t.Fatal(err)
	}

	samples, err := Samples(file)
	if err != nil {
		t.Fatal(err)
	}

	// the sample count should be close to:
	// sample rate * number of channels * duration
	want := 44100 * 2 * 213.574649
	if distance(len(samples), want) > 44100 {
		t.Errorf("invalid number of samples: %v, want %v", len(samples), want)
	}
}

func distance(a int, b float64) int {
	return int(math.Abs(float64(a) - b))
}
