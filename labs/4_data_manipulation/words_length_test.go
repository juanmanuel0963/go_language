package main

import (
	"reflect"
	"testing"
)

func TestGroupWordsByLengthWithCount(t *testing.T) {
	tests := []struct {
		input         []string
		expectedGroup map[int][]string
		expectedCount map[int]int
	}{
		{[]string{}, map[int][]string{}, map[int]int{}},
		{[]string{"cat", "dog", "bat"}, map[int][]string{3: {"cat", "dog", "bat"}}, map[int]int{3: 3}},
		{[]string{"hello"}, map[int][]string{5: {"hello"}}, map[int]int{5: 1}},
		{[]string{"apple", "banana", "apple"}, map[int][]string{5: {"apple", "apple"}, 6: {"banana"}}, map[int]int{5: 2, 6: 1}},
	}

	for _, test := range tests {
		resultGroup, resultCount := GroupWordsByLengthWithCount(test.input)
		if !reflect.DeepEqual(resultGroup, test.expectedGroup) || !reflect.DeepEqual(resultCount, test.expectedCount) {
			t.Errorf("Got %v and %v, expected %v and %v", resultGroup, resultCount, test.expectedGroup, test.expectedCount)
		}
	}
}
