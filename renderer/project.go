package renderer

import "math"

func max(points []int) int {
	m := math.MinInt64
	for _, p := range points {
		if p > m {
			m = p
		}
	}
	return m
}

// Project takes a slice of integers and newMax and returns a new slice
// of the same size, where all the values of the original slice are normalized
// to the ratio between newMax and the maximum number in the slice.
// For your convenience we provided a max function that retrieves the maximum element
// in a slice
// Example:
// given [1,2,6] and newMax 18
// the maximum element in the slice is 6 and the ratio between newMax and 6 is 3
// therefore the result slice will be [3,6,18]

func Project(in []int, newMax int) []int {
	return nil
}
