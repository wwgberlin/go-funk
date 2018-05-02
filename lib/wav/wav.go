package wav

import (
	"io"

	"github.com/go-audio/audio"
	_wav "github.com/go-audio/wav"
	"github.com/pkg/errors"
)

// Samples WAV samples and returns, if successful, a slice containing all
// the audio samples.
func Samples(r io.ReadSeeker) (samples []int, err error) {
	d := _wav.NewDecoder(r)
	if !d.IsValidFile() {
		return nil, errors.Errorf("invalid file")
	}

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
