package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceSwap2_v1(t *testing.T) {

	assert.Equal(t, SliceSwap2_v1([]int{1, 2, 3, 4, 5, 6}), []int{6, 5, 4, 3, 2, 1})

}

func TestSliceSwap2_v2(t *testing.T) {

	assert.Equal(t, SliceSwap2_v2([]int{1, 2, 3, 4, 5, 6}), []int{6, 5, 4, 3, 2, 1})

}
