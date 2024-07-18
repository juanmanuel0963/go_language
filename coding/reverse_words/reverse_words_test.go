package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{input: " the   sky   is   blue "},
			"blue is sky the",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.input); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution_1(t *testing.T) {

	got := reverseWords("the   sky   is   blue")
	assert.Equal(t, got, "blue is sky the")
}
