package programs

import (
	"fmt"
	"testing"
)

func TestSumOfSquares(t *testing.T) {

	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0

	go func() {

		fmt.Println("go func()")

		for i := 1; i <= 5; i++ {
			sum += <-mychannel

			toPrint := fmt.Sprintf("sum(Y^2): %v", sum)
			fmt.Println(toPrint)
		}

		quitchannel <- 0
	}()

	fmt.Println("intermedio")

	CalcSquares(mychannel, quitchannel)

}
