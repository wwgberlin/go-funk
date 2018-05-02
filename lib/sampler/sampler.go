package sampler

import (
	"errors"
	"fmt"
	"math"
)

func Sample(in []int, n int) ([]int, error) {
	if n <= 0 {
		return nil, fmt.Errorf("invalid sample count: %v", n)
	}
	if len(in) < n {
		return nil, errors.New("number of samples is bigger than the input")
	}
	if len(in) == n {
		return in, nil
	}

	samples := []int{}
	size := windowSize(len(in), n)

	var i int
	for ; i < n; i++ {
		value := 0
		low, high := i*size, (i+1)*size
		high = min(high, len(in))
		for _, v := range in[low:high] {
			value += v
		}
		samples = append(samples, value/(high-low))
	}

	return samples, nil
}

func windowSize(l, n int) int {
	return int(math.Ceil(float64(l) / float64(n)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
