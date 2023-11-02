package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		input1 [][]int
		intpu2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		/*{
			"Test 1",
			args{input1: [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}}, intpu2: []int{3, 4}},
			[][]int{{0, 1}, {2, 5}, {6, 7}, {8, 9}, {10, 11}},
		},*/
		{
			"Test 2",
			args{input1: [][]int{{1, 3}, {6, 9}}, intpu2: []int{2, 5}},
			[][]int{{1, 5}, {6, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := solution(tt.args.input1, tt.args.intpu2)

			for i := 0; i < len(got); i++ {

				for j := 0; j < len(got[i]); j++ {

					if got[i][j] != tt.want[i][j] {
						t.Errorf("solution() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func Test_solution_1(t *testing.T) {
	/*
		current := [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}}
		new := []int{3, 4}

		got := solution(current, new)

		assert.Equal(t, got, [][]int{{0, 1}, {2, 5}, {6, 7}, {8, 9}, {10, 11}})
	*/
	//-----------
	current1 := [][]int{{1, 3}, {6, 9}}
	new1 := []int{2, 5}

	got1 := solution(current1, new1)

	assert.Equal(t, got1, [][]int{{1, 5}, {6, 9}})
}
