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
		return nil, errors.New("count bigger than length")
	}
	size := windowSize(len(samples), count)
	compressed := make([]int, count)
	for s := 0; s < count; s++ {
		low, high := clamp(s*size, len(samples)), clamp((s+1)*size, len(samples))
		compressed[s] = fn(samples[low:high])
	}
	return compressed, nil
}

func windowSize(l, c int) int {
	return int(math.Ceil(float64(l) / float64(c)))
}

func clamp(v, max int) int {
	if v > max {
		return max
	}
	return v
}
