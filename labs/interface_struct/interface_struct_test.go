package InterfaceStruct

import (
	"fmt"
	"testing"
)

func TestInterfaceStruct(t *testing.T) {

	c1 := Circle{radius: 10}
	c1Area := CalculateArea(c1)

	if c1Area != 3.1416*c1.radius*c1.radius {
		exp := fmt.Sprintf("%f", 3.1416*c1.radius*c1.radius)
		t.Error("expected " + exp)
	}

	fmt.Println(c1Area)

	s1 := Square{sideA: 10, sideB: 10}
	s1Area := CalculateArea(s1)

	if s1Area != s1.sideA*s1.sideB {
		exp := fmt.Sprintf("%f", s1.sideA*s1.sideB)
		t.Error("expected " + exp)
	}

	fmt.Println(s1Area)
}
