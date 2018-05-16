package sampler

import (
	"math"
	"reflect"
	"testing"
)

func TestCompress(t *testing.T) {
	var acc []int
	var in []int

	fn := func(v []int) int {
		acc = append(acc, v...)
		return v[0]
	}

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
}

func TestCompressBounds(t *testing.T) {
	var acc []int
	var in []int

	fn := func(v []int) int {
		acc = append(acc, v...)
		sum := 0
		for _, i := range v {
			sum += i
		}
		return sum
	}

	in = []int{1, 2, 3, 4}
	s, err := Compress(in, 2, fn)
	if err != nil {
		t.Fatal("Compress was expected to succeed")
	}
	expected := []int{3, 7}
	if !reflect.DeepEqual(expected, s) {
		t.Fatalf("Compress was expected to return %v but returned %v", expected, s)
	}
	if !reflect.DeepEqual(in, acc) {
		t.Fatalf("Compress was not called with all the integers in the given slice expected: %v, got: %v", in, acc)
	}
}

func TestCompressBounds2(t *testing.T) {
	var acc []int
	var in []int

	fn := func(v []int) int {
		if len(v) == 0 {
			return math.MaxUint8
		}
		acc = append(acc, v...)
		sum := 0
		for _, i := range v {
			sum += i
		}
		return sum
	}

	in = []int{1, 2, 3, 4, 5, 6, 7, 8}
	s, err := Compress(in, 6, fn)
	if err != nil {
		t.Fatal("Compress was expected to succeed")
	}
	expected := []int{3, 7, 11, 15, 255, 255}
	if !reflect.DeepEqual(expected, s) {
		t.Fatalf("Compress was expected to return %v but returned %v", expected, s)
	}
	if !reflect.DeepEqual(in, acc) {
		t.Fatalf("Compress was not called with all the integers in the given slice expected: %v, got: %v", in, acc)
	}
}

func TestCompressError(t *testing.T) {
	var in []int

	fn := func(v []int) int {
		t.Fatal("sampler was not expected te be called")
		return 0
	}

	//Test transform with count larger than number of elements
	_, err := Compress(in, len(in)+1, fn)
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
