package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsPalindrome(t *testing.T) {

	isPalindrome1 := IsPalindrome("JUAN")
	assert.Equal(t, false, isPalindrome1)

	isPalindrome2 := IsPalindrome("TAT")
	assert.Equal(t, true, isPalindrome2)

}

func Test_solution(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test",
			args{inputString: "aabaa"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.inputString); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
