package miscellaneous

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExceptionHandling(t *testing.T) {

	f, err := os.Open("file.txt")

	if err != nil {
		log.Fatal(err)
	}

	if f != nil {
		fmt.Println("File name: ")
		fmt.Println(f.Name())
	}

	assert.Equal(t, "file.txt", f.Name())
}
