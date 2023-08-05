package channels

import (
	"fmt"
	"time"
)

func HelloWorld(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}
