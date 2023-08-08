package slices

type MyInterface interface {
	// Find number of elements in collection
	Len() int
	// Less method is used for identifying which elements among index i and j are l
	Less(i, j int) bool
	// Swap method is used for swapping elements with indexes i and j
	Swap(i, j int)
}

type Human struct {
	name string
	age  int
}

// AgeFactor implements sort.Interface that sorts the slice based on age field.
type AgeFactor []Human

func (a AgeFactor) Len() int { return len(a) }

func (a AgeFactor) Less(i, j int) bool { return a[i].age < a[j].age }
func (a AgeFactor) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
