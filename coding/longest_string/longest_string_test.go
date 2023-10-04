package main

import (
	"fmt"
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
			args{input: []string{"aba", "aa", "ad", "vcd", "aba"}},
			[]string{"aba", "vcd", "aba"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := solution(tt.args.input)

			for i := range got {
				if got[i] != tt.want[i] {
					fmt.Println("slices are not equal")
					t.Errorf("solution() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func Test_solution_1(t *testing.T) {

	got := solution([]string{"aba", "aa", "ad", "vcd", "aba"})
	assert.Equal(t, got, []string{"aba", "vcd", "aba"})

}
