package command_line

import (
	"fmt"
	"os"
	"testing"
)

// go run command_with_args.go twitter facebook whatsapp
// go run -race command_with_args.go twitter facebook whatsapp
// go test -race

func TestCommandWithArgs(t *testing.T) {

	fmt.Println("Init")

	toPrint := fmt.Sprintf("Args len: %v", len(os.Args))
	fmt.Println(toPrint)

	for i := 0; i < len(os.Args); i++ {
		toPrint := fmt.Sprintf("Args[%v]: %v", i, os.Args[i])
		fmt.Println(toPrint)
	}

}
