package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		rows         int
		cols         int
		participants [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test 1",
			args{rows: 3, cols: 5, participants: [][]int{{0, 0}, {0, 4}, {2, 2}}},
			[]int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := solution(tt.args.rows, tt.args.cols, tt.args.participants)

			for i := 0; i < len(got); i++ {

				if got[i] != tt.want[i] {
					t.Errorf("solution() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}

func Test_solution_1(t *testing.T) {

	got := solution(3, 5, [][]int{{0, 0}, {0, 4}, {2, 2}})
	assert.Equal(t, got, []int{0, 2})
}
