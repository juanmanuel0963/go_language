package main

import (
	"reflect"
	"testing"
)

func TestAdjustIngredients(t *testing.T) {
	type args struct {
		ingredients []string
		persons     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test1",
			args{[]string{"2 eggs", "200 grams of flour"}, 3},
			[]string{"6 eggs", "600 grams of flour"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdjustIngredients(tt.args.ingredients, tt.args.persons); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdjustIngredients() = %v, want %v", got, tt.want)
			}
		})
	}
}
