package slices

func SliceCompare(A_slice []int, B_slice []int) bool {

	if len(A_slice) != len(B_slice) {

		return false

	}

	for i := range A_slice {

		if A_slice[i] != B_slice[i] {

			return false

		}

	}

	return true
}
