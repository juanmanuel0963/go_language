package strings

import (
	"fmt"
	"testing"
)

func TestStringsMultiline(t *testing.T) {

	var Raw_String_Literals = `uninterrupted \n string`
	fmt.Println(Raw_String_Literals)

	var Interpreted_String_Literals = "interrupted \n string"
	fmt.Println(Interpreted_String_Literals)
}
