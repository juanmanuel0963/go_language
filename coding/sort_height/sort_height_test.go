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
		want []int
	}{
		{
			"Test 1",
			args{input: []int{-1, 150, 190, 170, -1, -1, 160, 180}},
			[]int{-1, 150, 160, 170, -1, -1, 180, 190},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := solution(tt.args.input)

			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("solution() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}

func Test_solution_1(t *testing.T) {

	got := solution([]int{-1, 150, 190, 170, -1, -1, 160, 180})
	assert.Equal(t, got, []int{-1, 150, 160, 170, -1, -1, 180, 190})
}
