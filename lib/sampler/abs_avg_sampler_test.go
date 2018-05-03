package sampler

import (
	"reflect"
	"testing"
)

func TestAbsAvg(t *testing.T) {
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
		{"success", args{[]int{1, -2, 255}, 3}, []int{1, 2, 255}, false},
		{"success", args{[]int{1, -3, 2, 4, 254, 256}, 3}, []int{2, 3, 255}, false},
		{"success", args{[]int{1, 3, 2, 4, 3, 5, 255}, 4}, []int{2, 3, 4, 255}, false},
		{"failure", args{[]int{1, 2, 3}, 4}, nil, true},
		{"failure", args{[]int{}, -42}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AbsAvg(tt.args.in, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("AbsAvgSample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AbsAvgSample() = %v, want %v", got, tt.want)
			}
		})
	}
}
