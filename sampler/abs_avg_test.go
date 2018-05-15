package sampler

import (
	"testing"
)

func TestAbsAvg(t *testing.T) {
	res := AbsAvg([]int{1, 3})
	if res != 2 {
		t.Errorf("AbsAvg expected to return 2 but returned %d", res)
	}

	res = AbsAvg([]int{3, -3})
	if res != 3 {
		t.Errorf("AbsAvg expected to return 3 but returned %d", res)
	}
}
