package functions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_swap_v1(t *testing.T) {

	x := "juan"
	y := "manuel"
	gotx, goty := swap(x, y)

	fmt.Println(gotx)
	fmt.Println(goty)

	assert.Equal(t, "manuel", gotx)
	assert.Equal(t, "juan", goty)

}

func Test_swap_v2(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			"test1",
			args{"juan", "manuel"},
			"manuel",
			"juan",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := swap(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("swap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("swap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
