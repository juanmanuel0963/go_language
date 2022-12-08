package main

import (
	"fmt"
	"os"
)

// go run command_with_args.go twitter facebook whatsapp

func main() {

	fmt.Println("Init")

	toPrint := fmt.Sprintf("Args len: %v", len(os.Args))
	fmt.Println(toPrint)

	for i := 0; i < len(os.Args); i++ {
		toPrint := fmt.Sprintf("Args[%v]: %v", i, os.Args[i])
		fmt.Println(toPrint)
	}

}
