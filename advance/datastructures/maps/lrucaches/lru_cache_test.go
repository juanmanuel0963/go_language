package lrucaches

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache_Contains(t *testing.T) {
	cache := New[string, int](5)
	cache.Put("hello", 1)

	assert.Equal(t, cache.Contains("hello"), true)
	assert.Equal(t, cache.Contains("unknown"), false)
	assert.Equal(t, cache.ContainsAll("hello", "unknown"), false)
	assert.Equal(t, cache.ContainsAny("hello", "unknown"), true)
}

func TestLRUCache_Clear(t *testing.T) {
	cache := New[string, int](5)
	assert.Equal(t, cache.IsEmpty(), true)
	assert.Equal(t, cache.Size(), 0)

	cache.Put("hello", 1)
	assert.Equal(t, cache.IsEmpty(), false)
	assert.Equal(t, cache.Size(), 1)

	cache.Clear()
	assert.Equal(t, cache.IsEmpty(), true)
	assert.Equal(t, cache.Size(), 0)
}

func TestLRUCache_GetKeys(t *testing.T) {
	cache := New[int, string](3)
	assert.True(t, reflect.DeepEqual(cache.GetKeys(), []int{}))
	assert.Equal(t, cache.Size(), 0)

	cache.Put(111, "a")
	assert.True(t, reflect.DeepEqual(cache.GetKeys(), []int{111}))
	assert.Equal(t, cache.Size(), 1)

	cache.Put(222, "b")
	assert.True(t, reflect.DeepEqual(cache.GetKeys(), []int{222, 111}))
	assert.Equal(t, cache.Size(), 2)

	cache.Put(333, "c")
	assert.True(t, reflect.DeepEqual(cache.GetKeys(), []int{333, 222, 111}))
	assert.Equal(t, cache.Size(), 3)

	cache.Put(444, "d")
	assert.True(t, reflect.DeepEqual(cache.GetKeys(), []int{444, 333, 222}))
	assert.Equal(t, cache.Size(), 3)

	cache.GetValue(333)
	assert.True(t, reflect.DeepEqual(cache.GetKeys(), []int{333, 444, 222}))
	assert.Equal(t, cache.Size(), 3)

	cache.Put(555, "f")
	assert.True(t, reflect.DeepEqual(cache.GetKeys(), []int{555, 333, 444}))
	assert.Equal(t, cache.Size(), 3)

	cache.Remove(333)
	assert.True(t, reflect.DeepEqual(cache.GetKeys(), []int{555, 444}))
	assert.Equal(t, cache.Size(), 2)
}

func TestLRUCache_GetEntries(t *testing.T) {
	cache := New[int, string](3)
	assert.True(t, reflect.DeepEqual(cache.GetEntries(), []*LRUEntry[int, string]{}))
	assert.Equal(t, cache.Size(), 0)

	cache.Put(111, "a")
	assert.True(t, reflect.DeepEqual(cache.GetEntries(), []*LRUEntry[int, string]{{Key: 111, Value: "a"}}))
	assert.Equal(t, cache.Size(), 1)

	cache.Put(222, "b")
	assert.True(t, reflect.DeepEqual(cache.GetEntries(), []*LRUEntry[int, string]{{Key: 222, Value: "b"}, {Key: 111, Value: "a"}}))
	assert.Equal(t, cache.Size(), 2)

	cache.Put(333, "c")
	assert.True(t, reflect.DeepEqual(cache.GetEntries(), []*LRUEntry[int, string]{{Key: 333, Value: "c"}, {Key: 222, Value: "b"}, {Key: 111, Value: "a"}}))
	assert.Equal(t, cache.Size(), 3)

	cache.Put(444, "d")
	assert.True(t, reflect.DeepEqual(cache.GetEntries(), []*LRUEntry[int, string]{{Key: 444, Value: "d"}, {Key: 333, Value: "c"}, {Key: 222, Value: "b"}}))
	assert.Equal(t, cache.Size(), 3)

	cache.GetValue(333)
	assert.True(t, reflect.DeepEqual(cache.GetEntries(), []*LRUEntry[int, string]{{Key: 333, Value: "c"}, {Key: 444, Value: "d"}, {Key: 222, Value: "b"}}))
	assert.Equal(t, cache.Size(), 3)

	cache.Put(555, "f")
	assert.True(t, reflect.DeepEqual(cache.GetEntries(), []*LRUEntry[int, string]{{Key: 555, Value: "f"}, {Key: 333, Value: "c"}, {Key: 444, Value: "d"}}))
	assert.Equal(t, cache.Size(), 3)

	cache.Remove(333)
	assert.True(t, reflect.DeepEqual(cache.GetEntries(), []*LRUEntry[int, string]{{Key: 555, Value: "f"}, {Key: 444, Value: "d"}}))
	assert.Equal(t, cache.Size(), 2)
}

func TestLRUCacheMerge(t *testing.T) {
	cache1 := New[bool, string](5)
	cache1.Put(true, "true")

	cache2 := New[bool, string](5)
	cache2.Put(true, "false")

	cache3 := New[bool, string](5)
	cache3.Merge(cache1, cache2)

	value, ok := cache3.GetValue(true)

	assert.Equal(t, cache3.Size(), 1)
	assert.True(t, ok)
	assert.Equal(t, value, "false")
	assert.Equal(t, cache3.ContainsAll(true), true)
}
