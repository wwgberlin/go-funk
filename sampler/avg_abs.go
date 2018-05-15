package sampler

import (
	"math"
)

// AvgAbs calculates the absolute value of the average
// of the values in the given slice
func AvgAbs(in []int) int {
	sum := 0
	for _, v := range in {
		sum += v
	}
	return int(math.Abs(float64(sum) / float64(len(in))))
}
