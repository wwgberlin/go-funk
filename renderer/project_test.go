package renderer

import (
	"reflect"
	"testing"
)

func TestProject(t *testing.T) {
	s := []int{1, 2, 6}
	e := []int{3, 6, 18}

	p := Project(s, 18)

	if !reflect.DeepEqual(p, e) {
		t.Fatalf("Unexpected result in Project. Expected %v but Got: %v", e, p)
	}
}
