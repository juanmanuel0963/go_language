package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
)

func TestSliceReverse(t *testing.T) {
	assert.Equal(t, SliceReverse_v1([]int{1, 2, 3, 4, 5, 6}), []int{6, 5, 4, 3, 2, 1})
	assert.NotEqual(t, SliceReverse_v2([]int{1, 2, 3, 4, 5, 6, 7}), []int{8, 7, 6, 5, 4, 3, 2, 1})
	assert.Equal(t, SliceReverse_v3([]int{1, 2, 3, 4, 5, 6, 7, 8}), []int{8, 7, 6, 5, 4, 3, 2, 1})
}
