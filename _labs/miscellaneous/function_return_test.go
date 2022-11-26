package miscellaneous

import "testing"

func Test_swap(t *testing.T) {
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
