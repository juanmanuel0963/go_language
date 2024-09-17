package heightChecker

import "testing"

func Test_heightChecker(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "Example 1",
			heights: []int{1, 1, 4, 2, 1, 3},
			want:    3,
		},
		{
			name:    "Example 2",
			heights: []int{5, 1, 2, 3, 4},
			want:    5,
		},
		{
			name:    "Example 3",
			heights: []int{1, 2, 3, 4, 5},
			want:    0,
		},
	}
	for _, tt := range tests {
		if got := heightChecker(tt.heights); got != tt.want {
			t.Errorf("%q. heightChecker() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_heightCheckerVer2(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "Example 1",
			heights: []int{1, 1, 4, 2, 1, 3},
			want:    3,
		},
		{
			name:    "Example 2",
			heights: []int{5, 1, 2, 3, 4},
			want:    5,
		},
		{
			name:    "Example 3",
			heights: []int{1, 2, 3, 4, 5},
			want:    0,
		},
	}
	for _, tt := range tests {
		if got := heightCheckerVer2(tt.heights); got != tt.want {
			t.Errorf("%q. heightChecker() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_heightCheckerVer3(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "Example 1",
			heights: []int{1, 1, 4, 2, 1, 3},
			want:    3,
		},
		{
			name:    "Example 2",
			heights: []int{5, 1, 2, 3, 4},
			want:    5,
		},
		{
			name:    "Example 3",
			heights: []int{1, 2, 3, 4, 5},
			want:    0,
		},
	}
	for _, tt := range tests {
		if got := heightCheckerVer3(tt.heights); got != tt.want {
			t.Errorf("%q. heightChecker() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Benchmark_heightChecker(b *testing.B) {
	heights := []int{1, 1, 4, 2, 1, 3}
	for i := 0; i < b.N; i++ {
		heightChecker(heights)
	}
}

func Benchmark_heightCheckerVer2(b *testing.B) {
	heights := []int{1, 1, 4, 2, 1, 3}
	for i := 0; i < b.N; i++ {
		heightCheckerVer2(heights)
	}
}

func Benchmark_heightCheckerVer3(b *testing.B) {
	heights := []int{1, 1, 4, 2, 1, 3}
	for i := 0; i < b.N; i++ {
		heightCheckerVer3(heights)
	}
}
