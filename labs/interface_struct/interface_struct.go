package InterfaceStruct

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	sideA float64
	sideB float64
}

func (c Circle) Area() float64 {
	return 3.1416 * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.1416 * c.radius
}

func (s Square) Area() float64 {
	return s.sideA * s.sideB
}
func (s Square) Perimeter() float64 {
	return 2*s.sideA + 2*s.sideB
}

func CalculateArea(sh Shape) float64 {
	return sh.Area()
}
