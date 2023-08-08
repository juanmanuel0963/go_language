package interfaces

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeasure_v1(t *testing.T) {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	areaR, perimR := Measure(r)
	fmt.Printf("area R: %0.2f", areaR)
	fmt.Println("")
	fmt.Printf("perim R: %0.2f", perimR)
	fmt.Println("")

	assert.Equal(t, areaR, float64(r.width*r.height))
	assert.Equal(t, perimR, float64(2*r.width+2*r.height))

	areaC, perimC := Measure(c)
	fmt.Printf("area C: %0.2f", areaC)
	fmt.Println("")
	fmt.Printf("perim C: %0.2f", perimC)
	fmt.Println("")

	assert.Equal(t, areaC, float64(math.Pi*c.radius*c.radius))
	assert.Equal(t, perimC, float64(2*math.Pi*c.radius))

}

func TestMeasure(t *testing.T) {
	type args struct {
		g geometry
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{
			"test1",
			args{rect{width: 3, height: 4}},
			3 * 4,
			2*3 + 2*4,
		},
		{
			"test2",
			args{circle{radius: 5}},
			float64(math.Pi * 5 * 5),
			float64(2 * math.Pi * 5),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Measure(tt.args.g)
			if got != tt.want {
				t.Errorf("Measure() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Measure() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
