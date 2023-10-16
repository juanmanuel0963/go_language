package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{input: []int{1, 3, 6, 4, 1, 2}},
			5,
		},
		{
			"Test 2",
			args{input: []int{1, 2, 3}},
			4,
		},
		{
			"Test 3",
			args{input: []int{-1, -3}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.input); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution_1(t *testing.T) {

	got := Solution([]int{1, 3, 6, 4, 1, 2})
	assert.Equal(t, got, 0)
}
