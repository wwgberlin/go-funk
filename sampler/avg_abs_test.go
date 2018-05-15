package sampler

import (
	"testing"
)

func TestAvgAbs(t *testing.T) {
	res := AvgAbs([]int{1, 3})
	if res != 2 {
		t.Errorf("AvgAbs expected to return 2 but returned %d", res)
	}

	res = AvgAbs([]int{3, -3})
	if res != 0 {
		t.Errorf("AvgAbs expected to return 0 but returned %d", res)
	}
}
