package strucs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructCompare_v1(t *testing.T) {

	person_a := Person{name: "Juan", age: 20}
	person_b := Person{name: "Juan", age: 20}
	person_c := Person{name: "Juan", age: 20}

	assert.Equal(t, StrucsCompare_Persons(person_a, person_b), true)
	assert.NotEqual(t, StrucsCompare_Persons(person_a, person_c), false)
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
			args{Person{name: "Juan", age: 20}, Person{name: "Juan", age: 20}},
			true,
		},
		{
			"test2",
			args{Person{name: "Juan", age: 20}, Person{name: "Manuel", age: 20}},
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
