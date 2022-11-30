package maps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMap_v1(t *testing.T) {

	search_map := map[int]string{100: "apple", 200: "pearl", 300: "pineapple"}
	search_key := 200

	val := SearchMap(search_map, search_key)
	assert.Equal(t, "pearl", val)

	printString := fmt.Sprintf("Search value is: %v", val)
	fmt.Println(printString)
}

func TestSearchMap_v2(t *testing.T) {
	type args struct {
		search_map map[int]string
		search_val int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{map[int]string{100: "apple", 200: "pearl", 300: "pineapple"}, 200},
			"pearl",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchMap(tt.args.search_map, tt.args.search_val); got != tt.want {
				t.Errorf("SearchMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
