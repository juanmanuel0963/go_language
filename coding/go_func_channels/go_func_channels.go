package main

import "fmt"

/*
func main() {
	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	println(counter) //0
}
*/
/*

The issue in this Go code is related to the use of goroutines and the way they access and modify the counter variable concurrently. Let me explain the problem:

Goroutines are being launched inside a loop:
You are launching 1000 goroutines inside a for loop. Each goroutine increments the counter variable. Since these goroutines are executing concurrently, there is no guarantee about the order in which they execute.

Race condition:
Multiple goroutines are accessing and modifying the counter variable concurrently without proper synchronization. This can lead to a race condition where multiple goroutines try to update counter at the same time, leading to unpredictable and incorrect results.

println(counter) is not synchronized:
Even though you are trying to print the counter variable outside the loop, there's no synchronization mechanism in place to ensure that all the goroutines have finished their work before printing the final value of counter. Therefore, you may see a value of 0 or some other unpredictable value when you print counter.

To fix this issue and get the correct count of how many goroutines have successfully incremented the counter variable, you can use channels for synchronization.
*/

func main() {
	channels()
}

func channels() {

	var counter int
	counter = 0

	for i := 0; i < 1000; i++ {

		myChannel := make(chan int)

		go func(channel chan int, myCounter int) {

			myCounter++
			channel <- myCounter
			close(channel)

		}(myChannel, counter)

		counter = <-myChannel
		fmt.Println(counter)
	}
	//fmt.Println(counter)
}

/*

The code you provided doesn't have a race condition because it effectively synchronizes access to the counter variable using channels. Let's break down how it works:

You initialize counter to 0 outside the loop.
Inside the loop, you create a new channel myChannel for each iteration.
You launch a goroutine for each iteration, passing the myChannel and the current value of counter (myCounter) as arguments.
Inside each goroutine, you increment myCounter, send it through the channel, and then close the channel.
After launching the goroutine, you immediately read from myChannel, which effectively waits until the goroutine sends a value on the channel.
You assign the value received from the channel to counter.
Finally, you print the updated value of counter.
This code ensures that each goroutine updates myCounter independently, sends its updated value through the channel, and then the main loop waits for the value to be received and updates the shared counter variable.

There's no race condition because you're using a separate channel (myChannel) for each goroutine to communicate the updated value of myCounter back to the main loop. This synchronization mechanism ensures that the updates to counter are sequential and don't lead to data races.

*/

/*
func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {

			defer wg.Done()
			mu.Lock()
			counter++
			fmt.Println(counter)
			mu.Unlock()
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	//fmt.Println(counter)
}
*/
