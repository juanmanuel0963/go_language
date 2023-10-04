package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{input: [][]int{{0, 1, 1, 2}, {0, 5, 0, 0}, {2, 0, 3, 3}}},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.input); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution_1(t *testing.T) {

	got := solution([][]int{{0, 1, 1, 2}, {0, 5, 0, 0}, {2, 0, 3, 3}})
	assert.Equal(t, got, 9)
}
