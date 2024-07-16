package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{word1: "abc", word2: "pqr"},
			"apbqcr",
		},
		{
			"Test 2",
			args{word1: "a", word2: "pqr"},
			"apqr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternately(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution_1(t *testing.T) {
	got := mergeAlternately("abc", "pqr")
	assert.Equal(t, got, "apbqcr")
}

func Test_solution_2(t *testing.T) {
	got := mergeAlternately("a", "pqr")
	assert.Equal(t, got, "apqr")
}
