package slices

import (
	"fmt"
	"sort"
	"testing"
)

// How can you sort a slice of custom structs

// We can sort slices of custom structs by using sort.Sort and sort.Stable functions.
// These methods sort any collection that implements sort.Interface interface that has
// Len(), Less() and Swap() methods as shown in the code below:

func TestSortSliceStructs(t *testing.T) {

	audience := []Human{
		{"Alice", 35},
		{"Bob", 45},
		{"James", 25},
	}

	sort.Sort(AgeFactor(audience))
	fmt.Println(audience)
}
