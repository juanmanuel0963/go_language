package goto_jump

import (
	"fmt"
	"testing"
)

func TestGoToJump(t *testing.T) {

	fmt.Println("a")
	goto FINISH //Transfers control to the labeled statement
	fmt.Println("b")
FINISH:
	fmt.Println("c")
}
