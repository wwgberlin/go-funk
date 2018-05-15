package sampler

import (
	"reflect"
	"testing"
)

func TestTransform(t *testing.T) {
	var acc []int
	var in []int

	fn := func(v []int) int {
		if len(v) == 0 {
			t.Fatal("sampler received an empty slice")
		}
		acc = append(acc, v...)
		return v[0]
	}

	//Test basic transform
	in = []int{1, 2, 3}
	s, err := Compress(in, 2, fn)
	if err != nil {
		t.Fatal("Compress was expected to succeed")
	}
	if !reflect.DeepEqual([]int{1, 3}, s) {
		t.Fatalf("Compress was expected to return [1,3] but returned %v", s)
	}
	if !reflect.DeepEqual(in, acc) {
		t.Fatalf("Compress was not called with all the integers in the given slice expected: %v, got: %v", in, acc)
	}
	acc = acc[:0]

	//Test transform with count larger than number of elements
	s, err = Compress(in, len(in)+1, fn)
	if err == nil {
		t.Fatal("Compress was expected to fail")
	}
}

func TestWindowSize(t *testing.T) {
	if windowSize(1, 3) != 1 {
		t.Error("window size was expected to return 1")
	}
	if windowSize(11, 5) != 3 {
		t.Error("window size was expected to return 3")
	}

}
