package sampler

import (
	"math"
)

type SamplerFunc func([]int, int) ([]int, error)

func windowSize(l, n int) int {
	return int(math.Ceil(float64(l) / float64(n)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func project(in []int, max, rangeMax int) []int {
	out := in[:0]
	for _, v := range in {
		out = append(out, v*max/rangeMax)
	}
	return out
}
