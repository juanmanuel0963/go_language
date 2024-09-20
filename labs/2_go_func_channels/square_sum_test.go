package main

import "testing"

func Test_judgeSquareSumVer3(t *testing.T) {
	tests := []struct {
		name string
		c    int
		want bool
	}{
		{
			name: "testcase 1",
			c:    5,
			want: true,
		},
		{
			name: "testcase 2",
			c:    3,
			want: false,
		},
		{
			name: "testcase 3",
			c:    4,
			want: true,
		},
		{
			name: "testcase 4",
			c:    2,
			want: true,
		},
		{
			name: "testcase 5",
			c:    1,
			want: true,
		},
		{
			name: "testcase 6",
			c:    0,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSumVer3(tt.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_judgeSquareSumVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		judgeSquareSumVer3(5)
	}
}
