package main

import (
	"fmt"
	"strconv"
)

func fractions_parser(fractions []string) []string {

	fmt.Println("input fractions: ", fractions)
	fmt.Println("--------")
	results := make([]string, 0)

	//Iterate the array
	for _, element := range fractions {

		//fmt.Println(element)

		//Iterate the whole operation
		for i := 0; i < len(element); i++ {

			leftTopVal := 0
			leftBottomVal := 0
			RightTopVal := 0
			RightBottomVal := 0

			//Find the +
			if element[i] == '+' {

				//Get left side
				leftSide := element[:i]
				//fmt.Println(leftSide)

				//Iterate the left side
				for l := 0; l < len(leftSide); l++ {

					//Find the /
					if leftSide[l] == '/' {
						//Get the top number
						leftTopVal, _ = strconv.Atoi(leftSide[:l])
						//fmt.Println(leftTopVal)

						//Get the bottom number
						leftBottomVal, _ = strconv.Atoi(leftSide[l+1:])
						//fmt.Println(leftBottomVal)
					}
				}

				//Get right side
				rightSide := element[i+1:]
				//fmt.Println(rightSide)

				//Iterate the right side
				for r := 0; r < len(rightSide); r++ {

					//Find the /
					if rightSide[r] == '/' {

						//Get the top number
						RightTopVal, _ = strconv.Atoi(rightSide[:r])
						//fmt.Println(RightTopVal)

						//Get the bottom number
						RightBottomVal, _ = strconv.Atoi(rightSide[r+1:])
						//fmt.Println(RightBottomVal)
					}
				}

				//fmt.Println(leftTopVal)
				//fmt.Println(leftBottomVal)
				//fmt.Println(RightTopVal)
				//fmt.Println(RightBottomVal)

				//Get numerator
				numerator := (leftTopVal * RightBottomVal) + (leftBottomVal * RightTopVal)

				//Get denominator
				denominator := (leftBottomVal * RightBottomVal)

				fmt.Println("numerator: ", numerator)
				fmt.Println("denominator: ", denominator)

				for {

					newNumerator, newDenominator := simplify(numerator, denominator)

					if newNumerator != numerator && newDenominator != denominator && newNumerator > 0 && newDenominator > 0 {

						numerator = newNumerator
						denominator = newDenominator

					} else {

						fmt.Println("numerator simplifyied: ", numerator)
						fmt.Println("denominator simplifyied: ", denominator)

						results = append(results, fmt.Sprint(numerator)+"/"+fmt.Sprint(denominator))
						//fmt.Println(results)
						break
					}
				}

			}

		}

		fmt.Println("--------")
	}

	return results
}

func simplify(numerator int, denominator int) (int, int) {

	minNumber := min(numerator, denominator)
	//fmt.Println("min number:", minNumber)

	newNumerator := 0
	newDenominator := 0

	//Iterate until min(numerator,denominator)
	for m := 2; m <= minNumber; m++ {

		modNumerator := numerator % m
		modDenominator := denominator % m

		//Find if mod(numerator/m)==0 and mod(denominator/m)==0
		if modNumerator == 0 && modDenominator == 0 {

			newNumerator = numerator / m
			newDenominator = denominator / m

			break
		}

	}

	return newNumerator, newDenominator
}

func main() {

	fmt.Println("Hey")

	fractions := []string{"2/3+4/6", "10/5+5/5"}
	result := fractions_parser(fractions)
	fmt.Println("Result: ", result)
}
