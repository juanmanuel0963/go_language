package main

import (
	"testing"
)

func TestComputeClosestToZero(t *testing.T) {
	type args struct {
		ts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]int{1, 2, 3, 4, 5}},
			1,
		},
		{
			"test2",
			args{[]int{-10, 20, 5, -20, 100}},
			5,
		},
		{
			"test3",
			args{[]int{-5, -4, -2, 12, -40, 4, 2, 18, 11, 5}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeClosestToZero(tt.args.ts); got != tt.want {
				t.Errorf("ComputeClosestToZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
