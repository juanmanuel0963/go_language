package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupWordsByLengthWithCountStudy(t *testing.T) {

	type TestCase struct {
		input         []string
		expectedGroup map[int][]string
		expectedCount map[int]int
	}

	AllTests := make([]TestCase, 0)

	case1 := TestCase{
		input:         []string{"apple", "bannana", "pineapple", "bannana"},
		expectedGroup: map[int][]string{5: {"apple"}, 7: {"bannana", "bannana"}, 9: {"pineapple"}},
		expectedCount: map[int]int{5: 1, 7: 2, 9: 1},
	}

	AllTests = append(AllTests, case1)

	for _, test := range AllTests {

		fmt.Println(test)

		resultGroup, resultCount, err := GroupWordsByLengthWithCountStudy(test.input)

		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(resultGroup, test.expectedGroup) || !reflect.DeepEqual(resultCount, test.expectedCount) {
			t.Errorf("Got %v and %v, expected %v and %v", resultGroup, resultCount, test.expectedGroup, test.expectedCount)
		}
	}
}
