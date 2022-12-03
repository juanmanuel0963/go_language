package maps

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapCopy_v1(t *testing.T) {

	map_from := map[string]bool{"A": true, "B": true}
	map_to := make(map[string]bool)
	map_want := map[string]bool{"A": true, "B": true}

	assert.False(t, reflect.DeepEqual(map_from, map_to))

	assert.True(t, reflect.DeepEqual(map_from, map_want))

	map_to, err := MapCopy(map_from, map_to)

	assert.Equal(t, map_want, map_to)

	assert.Nil(t, err)

	assert.True(t, reflect.DeepEqual(map_from, map_to))

	assert.NotEqual(t, map[string]bool{"A": false, "B": false}, map_to)

}

func TestMapCopy_v2(t *testing.T) {
	type args struct {
		from_map map[string]bool
		to_map   map[string]bool
	}
	tests := []struct {
		name    string
		args    args
		wantR   map[string]bool
		wantErr bool
	}{
		{
			"Test1",
			args{map[string]bool{"A": true, "B": true}, make(map[string]bool)},
			map[string]bool{"A": true, "B": true},
			false,
		},
		{
			"Test2",
			args{map[string]bool{"X": true, "Y": true}, make(map[string]bool)},
			map[string]bool{"X": true, "Y": true},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := MapCopy(tt.args.from_map, tt.args.to_map)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapCopy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("MapCopy() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestMapCopy_v3(t *testing.T) {

	map1 := map[string]bool{"Interview": true, "Bit": true}
	map2 := map[string]bool{"Interview": true, "Questions": true}
	map3 := map1
	map1 = map2 //copy description
	fmt.Println(map1, map2, map3)

	assert.Equal(t, map[string]bool{"Interview": true, "Questions": true}, map1)
	assert.Equal(t, map[string]bool{"Interview": true, "Questions": true}, map2)
	assert.Equal(t, map[string]bool{"Interview": true, "Bit": true}, map3)
}
