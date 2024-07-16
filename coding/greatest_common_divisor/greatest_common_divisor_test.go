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
			args{word1: "ABCABC", word2: "ABC"},
			"ABC",
		},
		{
			"Test 2",
			args{word1: "ABABAB", word2: "ABAB"},
			"AB",
		},
		{
			"Test 3",
			args{word1: "ABCDEF", word2: "ABC"},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution_1(t *testing.T) {
	got := gcdOfStrings("ABCABC", "ABC")
	assert.Equal(t, got, "ABC")
}

func Test_solution_2(t *testing.T) {
	got := gcdOfStrings("ABABAB", "ABAB")
	assert.Equal(t, got, "AB")
}
func Test_solution_3(t *testing.T) {
	got := gcdOfStrings("LEET", "CODE")
	assert.Equal(t, got, "")
}
func Test_solution_4(t *testing.T) {
	got := gcdOfStrings("ABCDEF", "ABC")
	assert.Equal(t, got, "")
}
func Test_solution_5(t *testing.T) {
	got := gcdOfStrings("NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM", "NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM")
	assert.Equal(t, got, "NLZGM")
}
