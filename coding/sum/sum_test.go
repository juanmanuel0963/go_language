package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_solution(t *testing.T) {

	got := solution(100, 50)

	assert.Equal(t, 150, got)

}

func Test_solution_1(t *testing.T) {
	type args struct {
		param1 int
		param2 int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{100, 50},
			150,
		},
		{
			"Test 2",
			args{200, 100},
			300,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.param1, tt.args.param2); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_solution_2(t *testing.T) {

	type args struct {
		param1 int
		param2 int
	}

	type test struct {
		name string
		args args
		want int
	}

	myTest := test{name: "Test 2", args: args{100, 50}, want: 150}

	got := solution(myTest.args.param1, myTest.args.param2)

	assert.Equal(t, 150, got)
}

func Test_solution_3(t *testing.T) {
	type args struct {
		param1 int
		param2 int
	}

	type test struct {
		name string
		args args
		want int
	}

	var myTest test
	myTest.name = "Test 3"
	myTest.args.param1 = 100
	myTest.args.param2 = 50
	myTest.want = 150

	got := solution(myTest.args.param1, myTest.args.param2)

	assert.Equal(t, 150, got)
}
