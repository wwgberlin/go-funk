package sampler

// AbsAvg calculates the average of the absolute values in the given slice
func AbsAvg(in []int) int {
	sum := 0
	for _, n := range in {
		if n < 0 {
			n = -n
		}
		sum += n
	}
	return sum / len(in)
}
