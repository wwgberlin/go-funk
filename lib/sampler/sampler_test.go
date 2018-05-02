package sampler

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	type args struct {
		in []int
		n  int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"success", args{[]int{1, 2, 3}, 3}, []int{1, 2, 3}, false},
		{"success", args{[]int{1, 3, 2, 4, 3, 5}, 3}, []int{2, 3, 4}, false},
		{"success", args{[]int{1, 3, 2, 4, 3, 5, 5}, 4}, []int{2, 3, 4, 5}, false},
		{"failure", args{[]int{1, 2, 3}, 4}, nil, true},
		{"failure", args{[]int{}, -42}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sample(tt.args.in, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sample() = %v, want %v", got, tt.want)
			}
		})
	}
}
