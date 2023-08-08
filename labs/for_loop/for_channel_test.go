package for_loop

import (
	"fmt"
	"testing"
	"time"
)

//for [condition | ( init; condition; increment ) | Range]
//{
//	statement(s);
//	more statements
//}

func TestForChannel(t *testing.T) {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
