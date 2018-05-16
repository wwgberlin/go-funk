package sampler

// AvgAbs calculates the absolute value of the average
// of the values in the given slice
func AvgAbs(in []int) int {
	sum := 0
	for _, n := range in {
		sum += n
	}
	avg := sum / len(in)
	if avg < 0 {
		avg = -avg
	}
	return avg
}
