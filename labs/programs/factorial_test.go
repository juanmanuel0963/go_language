package programs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial_v1(t *testing.T) {

	n := 3
	factorial := Factorial(n)
	fmt.Println(factorial)

	assert.Equal(t, 1, Factorial(1))
	assert.Equal(t, 6, Factorial(3))
	assert.Equal(t, 120, Factorial(5))
}

func TestFactorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{0},
			1,
		},
		{
			"test1",
			args{3},
			6,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
