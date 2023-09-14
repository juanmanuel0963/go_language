// CODE EXAMPLE VALID FOR COMPILING
package main

import "fmt"

type CustomErrors interface {
	Error() string
}

type MyError struct {
	Msg string
}

func getError(i CustomErrors) string {
	return i.Error()
}

func (s MyError) Error() string {
	return s.Msg
}

func main() {

	var ce1 MyError
	ce1.Msg = "Error #1"
	val1 := ce1.Error()
	fmt.Println(val1)

	ce2 := MyError{Msg: "Error #2"}
	val2 := getError(ce2)
	fmt.Println(val2)
}
