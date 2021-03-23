package _435

import "testing"

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{intervals: [][]int{
			{1, 2},
			{2, 3},
			{3, 4},
			{1, 3},
		}}, 1},
		{"t2", args{intervals: [][]int{
			{1, 2},
			{1, 2},
			{1, 2},
		}}, 2},
		{"t3", args{intervals: [][]int{
			{1, 2},
			{2, 3},
		}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
