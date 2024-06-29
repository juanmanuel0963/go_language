package files

import (
	"fmt"
	"log"
	"os"
)

//Explain the defer statement in Golang. Give an example of a deferred function’s call.
//-------------------------------------------------------------------------------------

//- A defer statement defers or postpones the execution of a function.
//- It postpones the execution until the surrounding function returns, either normally or through a panic call.
//- We use defer to ensure that a function call is performed later in the program’s execution, usually for cleaning resources.
//- For example, let’s say that we want to create a file, write to it, and then close when we’re done with it.
//- Immediately after creating the file variable, we defer the closing of that file.
//- The function that closes the file will be executed at the end of the enclosing function (main)
//  after the operation of writing to the file has finished.

func Defer() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer file.Close()
}
