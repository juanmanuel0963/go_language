package main

import "testing"

func Test_sumUntilEqual(t *testing.T) {
	type args struct {
		S1 int
		S2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{11, 7},
			620,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumUntilEqual(tt.args.S1, tt.args.S2); got != tt.want {
				t.Errorf("sumUntilEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
