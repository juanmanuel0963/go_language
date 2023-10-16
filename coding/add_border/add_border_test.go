package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Test 1",
			args{input: []string{"abc", "def"}},
			[]string{"*****", "*abc*", "*def*", "*****"},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := solution(tt.args.input)

			for i := 0; i < len(got); i++ {

				assert.Equal(t, got[i], tt.want[i])

				if got[i] != tt.want[i] {
					t.Errorf("solution() = %v, want %v", got[i], tt.want[i])
				}
			}

		})
	}
}
func Test_solution_1(t *testing.T) {

	got := solution([]string{"abc", "def"})
	assert.Equal(t, got, []string{"*****", "*abc*", "*def*", "*****"})
}
