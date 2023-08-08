package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructCompare_v1(t *testing.T) {

	person_a := Person{name: "Juan", age: 20, power: 100}
	person_b := Person{name: "Juan", age: 20, power: 100}
	person_c := Person{name: "Manuel", age: 20, power: 100}

	assert.Equal(t, true, person_a == person_b)
	assert.NotEqual(t, true, person_a == person_c)
}

func TestStrucsCompare_v2(t *testing.T) {
	type args struct {
		a Person
		b Person
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{Person{name: "Juan", age: 20, power: 100}, Person{name: "Juan", age: 20, power: 100}},
			true,
		},
		{
			"test2",
			args{Person{name: "Juan", age: 20, power: 100}, Person{name: "Manuel", age: 20, power: 100}},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrucsCompare_Persons(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("StrucsCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
// Ensure that they don’t contain slices, maps, or functions. Otherwise, the code won’t comply.
func TestStructCompare_v3(t *testing.T) {
	type Bar struct{ A []int }

	a := Bar{A: []int{1}}
	b := Bar{A: []int{1}}

	fmt.Println(a == b)
}
*/
// Ensure that they don’t contain slices, maps, or functions. Otherwise, the code won’t comply.
func TestStructCompare_v4(t *testing.T) {

	//animal_a := Animal{name: "Perro", size: 80, sons: []string{"Jacob"}}
	//animal_b := Animal{name: "Gato", size: 50, sons: []string{"Gora"}}

	animal_a := Animal{name: "Perro", size: 80}
	animal_b := Animal{name: "Gato", size: 50}

	assert.Equal(t, false, animal_a == animal_b)

}
