package utils

import (
	"math"
)

type SliceStats struct {
	Min int
	Max int
}

func IntSliceStats(in []int) *SliceStats {
	min, max := math.MaxInt64, math.MinInt64
	for _, v := range in {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return &SliceStats{
		Min: min,
		Max: max,
	}
}
