package main

/*
func Test_solution(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{input: 0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.input); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution_1(t *testing.T) {

	got := solution(0)
	assert.Equal(t, got, 0)
}
*/
/*
func recursiveSolution(grid [N][N]int, ch chan [N][N]int) {
	pending, row, col := pendingPositions(grid)

	if !pending && CheckBoard(grid) {
		ch <- grid
		close(ch)
		return
	}

	for i := 1; i <= 9; i++ {
		if valid(grid, i, row, col) {
			grid[row][col] = i
			go recursiveSolution(grid, ch)
		}
	}
}
*/
