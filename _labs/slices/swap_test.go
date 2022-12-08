package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
)

func TestSliceSwap(t *testing.T) {

	assert.Equal(t, SliceSwap_v0([]int{100, 200}), []int{200, 100})
	assert.Equal(t, SliceSwap_v1([]int{1, 2, 3, 4, 5, 6}), []int{6, 5, 4, 3, 2, 1})
	assert.NotEqual(t, SliceSwap_v2([]int{1, 2, 3, 4, 5, 6, 7}), []int{8, 7, 6, 5, 4, 3, 2, 1})
	assert.Equal(t, SliceSwap_v3([]int{1, 2, 3, 4, 5, 6, 7, 8}), []int{8, 7, 6, 5, 4, 3, 2, 1})
}
