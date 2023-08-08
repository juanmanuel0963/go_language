package channels

func SendReceive(n int, ch chan int) {

	//sending into the channel
	ch <- n * 2
}
