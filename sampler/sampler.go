package sampler

type SamplerFunc func([]int) int

var Samplers = map[string]SamplerFunc{
	"avg":     AvgAbs,
	"abs_avg": AbsAvg,
}

func RegSampler(name string, fn SamplerFunc) {
	Samplers[name] = fn
}
