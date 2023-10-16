package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 1",
			args{[]int{1, 2, 3}, []int{1, 2, 3}},
			true,
		},
		{
			"Test 2",
			args{[]int{1, 2, 3}, []int{2, 1, 3}},
			true,
		},
		{
			"Test 3",
			args{[]int{1, 2, 2}, []int{2, 1, 1}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution_1(t *testing.T) {

	got1 := solution([]int{1, 2, 3}, []int{1, 2, 3})
	assert.Equal(t, got1, true)

	got2 := solution([]int{1, 2, 3}, []int{2, 1, 3})
	assert.Equal(t, got2, true)

	got3 := solution([]int{1, 2, 2}, []int{2, 1, 1})
	assert.Equal(t, got3, false)

}
