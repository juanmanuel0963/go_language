package loops

import (
	"fmt"
	"testing"
)

//for [condition | ( init; condition; increment ) | Range]
//{
//	statement(s);
//	more statements
//}

func TestFor(t *testing.T) {
	// For loop to print numbers from 1 to 5
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}
}
