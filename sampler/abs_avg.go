package sampler

import "math"

// AbsAvg calculates the average of the absolute values in the given slice
func AbsAvg(in []int) int {
	sum := 0
	for _, i := range in {
		sum += int(math.Abs(float64(i)))
	}
	return sum / len(in)
}
