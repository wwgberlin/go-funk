package sampler

import (
	"errors"
	"math"
)

var errCountInvalid = errors.New("count must be <= length of the slice")

// Compress reduces the given slice to a new slice of length count using fn.
// Example:
// Given samples=[0,1,2,3,4], count=3, fn=sum
// We divide the samples into 3 windows: [0,1] [2,3] [4]
// call fn with each of the slices and add the result to the new slice
// The result slice will then be [1, 5. 4]
// * hint: samples[0:2], samples[2:4], samples[4:]
// * hint: use windowSize

func Compress(samples []int, count int, fn SamplerFunc) ([]int, error) {
	if count > len(samples) {
		return nil, errCountInvalid
	}
	out := make([]int, 0, count)
	ws := windowSize(len(samples), count)
	for w := 0; w < count; w++ {
		offset := ws * w
		size := ws
		if offset > len(samples) {
			out = append(out, fn(nil))
		} else if offset+size > len(samples) {
			out = append(out, fn(samples[offset:]))
		} else {
			out = append(out, fn(samples[offset:offset+size]))
		}
	}
	return out, nil
}

func windowSize(l, c int) int {
	return int(math.Ceil(float64(l) / float64(c)))
}
