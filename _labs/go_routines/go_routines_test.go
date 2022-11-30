package go_routines

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutines(t *testing.T) {

	//We see that the sampleRoutine() function is called by specifying the keyword go before it
	go sampleRoutine()

	//When a function is called a goroutine, the call will be returned
	//immediately to the next line of the program statement which is why “Started Main”
	//would be printed first and the goroutine will be scheduled and run concurrently in the background.
	fmt.Println("Started Main")

	//The sleep statement ensures that the goroutine is scheduled before
	//the completion of the main goroutine.
	time.Sleep(1 * time.Second)

	fmt.Println("Finished Main")

	//The output of this code would be:
	//Started Main
	//Inside Sample Goroutine
	//Finished Main
}
