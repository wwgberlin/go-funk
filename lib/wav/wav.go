package wav

import (
	"os"

	"github.com/go-audio/audio"
	_wav "github.com/go-audio/wav"
	"github.com/pkg/errors"
)

// Samples loads a wav file and returns, if successful, a slice containing all
// the audio samples.
func Samples(file string) ([]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d := _wav.NewDecoder(f)
	if !d.IsValidFile() {
		return nil, errors.Errorf("invalid file: %q", file)
	}

	samples := []int{}
	intBuf := make([]int, 256)
	buf := &audio.IntBuffer{Data: intBuf}

	for err == nil {
		n, err := d.PCMBuffer(buf)
		if err != nil {
			return nil, err
		}

		if n == 0 {
			break
		}

		samples = append(samples, buf.Data...)
	}

	return samples, nil
}
