package slices

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy_v1(t *testing.T) {
	assert.Equal(t, SliceCopy([]int{1, 2, 3, 4, 5}, make([]int, 5)), []int{1, 2, 3, 4, 5})
}

func TestSliceCopy(t *testing.T) {
	type args struct {
		from_slice []int
		to_slice   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test1",
			args{[]int{1, 2, 3, 4, 5}, make([]int, 5)},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"Test2",
			args{[]int{1, 2, 3, 4, 5, 6}, make([]int, 6)},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceCopy(tt.args.from_slice, tt.args.to_slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
