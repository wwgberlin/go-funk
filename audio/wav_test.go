package audio

import (
	"math"
	"os"
	"path/filepath"
	"testing"
)

func TestNewWavData(t *testing.T) {
	path, err := filepath.Abs("../public/rick/fixture.wav")
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	data, err := NewWavData(file)
	if err != nil {
		t.Fatal(err)
	}

	// the sample count should be close to:
	// sample rate * number of channels * duration
	want := 44100 * 2 * 213.574649
	if distance(len(data.Samples), want) > 44100 {
		t.Errorf("invalid number of samples: %v, want %v", len(data.Samples), want)
	}
}

func distance(a int, b float64) int {
	return int(math.Abs(float64(a) - b))
}
