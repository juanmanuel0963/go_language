package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			"Test 1",
			args{candies: []int{2, 3, 5, 1, 3}, extraCandies: 3},
			[]bool{true, true, true, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := kidsWithCandies(tt.args.candies, tt.args.extraCandies)

			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("solution() = %v, want %v", got, tt.want)
				}
			}
		})
	}

}

func Test_solution_1(t *testing.T) {

	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	got := kidsWithCandies(candies, extraCandies)

	assert.Equal(t, got, []bool{true, true, true, false, true})
}
