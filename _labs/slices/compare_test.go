package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceCompare_v1(t *testing.T) {

	A_slice := []int{10, 20, 30}
	B_slice := []int{10, 20, 30}
	assert.Equal(t, true, SliceCompare(A_slice, B_slice))

}

func TestSliceCompare_v2(t *testing.T) {
	type args struct {
		A_slice []int
		B_slice []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test1",
			args{[]int{10, 20, 30}, []int{10, 20, 30}},
			true,
		},
		{
			"Test1",
			args{[]int{10, 20, 30}, []int{10, 20, 30}},
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceCompare(tt.args.A_slice, tt.args.B_slice); got != tt.want {
				t.Errorf("SliceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
