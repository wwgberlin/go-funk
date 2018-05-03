package sampler

import (
	"errors"
	"fmt"
	"math"
)

func Avg(in []int, n int) ([]int, error) {
	if n <= 0 {
		return nil, fmt.Errorf("invalid sample count: %v", n)
	}
	if len(in) < n {
		return nil, errors.New("number of samples is bigger than the input")
	}

	samples := []int{}
	size := windowSize(len(in), n)

	max := math.MinInt64
	var i int
	for ; i < n; i++ {
		value := 0
		low, high := i*size, (i+1)*size
		high = min(high, len(in))
		for _, v := range in[low:high] {
			value += v
		}
		sample := int(math.Abs(float64(value / (high - low))))
		if sample > max {
			max = sample
		}
		samples = append(samples, sample)
	}

	return project(samples, 255, max), nil
}
