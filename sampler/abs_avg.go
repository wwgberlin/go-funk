package sampler

// AbsAvg calculates the average of the absolute values in the given slice
func AbsAvg(in []int) int {
	acc := 0
	for _, v := range in {
		acc += abs(v)
	}
	return acc / len(in)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
