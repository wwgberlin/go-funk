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
	points := make([]int, count)
	window := windowSize(len(samples), count)
	for i := 0; i < count; i++ {
		low := i * window
		high := (i + 1) * window
		if high > len(samples)-1 {
			points[i] = fn(samples[low:])
		} else {
			points[i] = fn(samples[low:high])
		}
	}
	return points, nil
}

func windowSize(l, c int) int {
	return int(math.Ceil(float64(l) / float64(c)))
}
