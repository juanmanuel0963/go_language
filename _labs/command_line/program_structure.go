// package clause
// It starts every Go source file.
// main is a special name declaring an executable rather than a library (package)

package main1

// import declaration.
// It declares packages used in this file.
// fmt is a standard library package uses to print to standard output (console)
import "fmt"

// function declaration.
// main is a special function name. it’s the entry point for the executable program

func main1() {
	// local scoped variables and constants declarations, calling functions etc
	var age int = 20
	var today string = "Monday"

	// Println() function from fmt package prints out a line to stdout
	fmt.Println("Today is", today)  // => Today is Monday
	fmt.Println("Your age is", age) // => You age is 17
}
