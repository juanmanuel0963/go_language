package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{input: 1},
			1,
		},
		{
			"Test 2",
			args{input: 2},
			5,
		},
		{
			"Test 3",
			args{input: 3},
			13,
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

	got := solution(1)
	assert.Equal(t, got, 1)
}
