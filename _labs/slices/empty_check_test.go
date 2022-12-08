package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {

	yearsSlice := []int{2000, 2010, 2020, 2030}

	isEmptyResponse := IsEmpty(yearsSlice)

	toPrint := fmt.Sprintf("Slice empty: %v", isEmptyResponse)

	fmt.Println(toPrint)

	assert.Equal(t, false, isEmptyResponse)
}
