package groupanagrams

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	values := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	multimap := GroupAnagrams(values)

	assert.Equal(t, multimap.Size(), 3)
	assert.True(t, reflect.DeepEqual(multimap.GetValues("aet"), []string{"eat", "tea", "ate"}))
	assert.True(t, reflect.DeepEqual(multimap.GetValues("ant"), []string{"tan", "nat"}))
	assert.True(t, reflect.DeepEqual(multimap.GetValues("abt"), []string{"bat"}))
}
