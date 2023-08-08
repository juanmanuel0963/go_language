package structs

import "fmt"

var numOfRedCars = 0

type Car struct {
	name  string
	model string
	color string
}

func StructSlices(MyCars []Car) int {

	for i := 0; i < len(MyCars); i++ {
		if MyCars[i].color == "red" {
			numOfRedCars += 1
			fmt.Println(MyCars[i].color)
			fmt.Println("Number of Red Cars ", numOfRedCars)
		}
	}

	return numOfRedCars
}
