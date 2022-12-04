package structs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructSlices_v1(t *testing.T) {

	MyCars := []Car{{name: "Toyota", model: "Corolla", color: "red"}, {name: "Toyota", model: "Innova", color: "gray"}}

	numOfRedCars := StructSlices(MyCars)
	fmt.Println(numOfRedCars)

	assert.Equal(t, 1, numOfRedCars)

}

func TestStructSlices(t *testing.T) {
	type args struct {
		MyCars []Car
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]Car{{name: "Toyota", model: "Corolla", color: "red"}, {name: "Toyota", model: "Innova", color: "gray"}}},
			1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructSlices(tt.args.MyCars); got != tt.want {
				t.Errorf("StructSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
