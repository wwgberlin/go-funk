package audio

import (
	"io"
	"time"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

type WavData struct {
	Duration time.Duration
	Samples  []int
}

// Samples WAV samples and returns, if successful, a slice containing all
// the audio samples.
func NewWavData(r io.ReadSeeker) (data WavData, err error) {
	d := wav.NewDecoder(r)

	intBuf := make([]int, 256)
	buf := audio.IntBuffer{Data: intBuf}

	for err == nil {
		n, err := d.PCMBuffer(&buf)
		if err != nil {
			return data, err
		}

		if n == 0 {
			break
		}

		data.Samples = append(data.Samples, buf.Data...)
	}

	data.Duration, err = d.Duration()
	return
}
