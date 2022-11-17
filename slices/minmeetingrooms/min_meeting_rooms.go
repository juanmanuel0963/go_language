package minmeetingrooms

import "github.com/juanmanuel0963/go_language/v1/datastructures/sets/hashmultisets"

// MinMeetingRooms returns minimum meeting rooms required
func MinMeetingRooms(intervals [][]int) int {
	multiset := hashmultisets.New[int]()
	for _, interval := range intervals {
		for i := interval[0]; i <= interval[1]; i++ {
			multiset.Add(i)
		}
	}

	multiSetPairs := multiset.GetTopValues()
	if len(multiSetPairs) == 0 {
		return 0
	}
	return multiSetPairs[0].Count
}
