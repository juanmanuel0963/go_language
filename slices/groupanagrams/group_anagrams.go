package groupanagrams

import (
	"github.com/juanmanuel0963/go_language/v1/datastructures/maps/hashmultimaps"
	"github.com/juanmanuel0963/go_language/v1/strings/sorts"
)

// GroupAnagrams groups anagrams
func GroupAnagrams(list []string) *hashmultimaps.HashMultiMap[string, string] {
	multimap := hashmultimaps.New[string, string]()
	for _, value := range list {
		sortedValue := sorts.Sort(value)
		multimap.Put(sortedValue, value)
	}
	return multimap
}
