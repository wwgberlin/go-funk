package sampler

// AvgAbs calculates the absolute value of the average
// of the values in the given slice
func AvgAbs(in []int) int {
	acc := 0
	for _, v := range in {
		acc += v
	}
	return acc / len(in)
}
