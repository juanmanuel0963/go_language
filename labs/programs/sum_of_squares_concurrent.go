package programs

import "fmt"

func CalcSquares(mychannel chan int, quit chan int) {

	fmt.Println("CalcSquares()")

	y := 1
	for {
		select {
		case mychannel <- (y * y):
			toPrint := fmt.Sprintf("Y^2: %v", y*y)
			fmt.Println(toPrint)
			y++
		case <-quit:
			return
		}
	}
}
