package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{year: 2023},
			21,
		},
		{
			"Test 2",
			args{year: 2000},
			20,
		},
		{
			"Test 3",
			args{year: 1978},
			20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.year); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution_1(t *testing.T) {
	got := solution(2005)
	assert.Equal(t, got, 21)
}
