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
			args{input: "(bar)"},
			"rab",
		},
		{
			"Test 2",
			args{input: "foo(bar)baz"},
			"foorabbaz",
		},
		{
			"Test 3",
			args{input: "foo(bar)baz(blim)"},
			"foorabbazmilb",
		},
		{
			"Test 4",
			args{input: "foo(bar(baz))blim"},
			"foobazrabblim",
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

	got1 := solution("(bar)")
	assert.Equal(t, got1, "rab")

	got2 := solution("foo(bar)baz")
	assert.Equal(t, got2, "foorabbaz")

	got3 := solution("foo(bar)baz(blim)")
	assert.Equal(t, got3, "foorabbazmilb")

	got4 := solution("foo(bar(baz))blim")
	assert.Equal(t, got4, "foobazrabblim")
}
