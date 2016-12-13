package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("Channels Started!")

	done := make(chan bool, 1)
	error := make(chan bool, 1)
	go func() {
		fmt.Println("Done GoROutine")
		time.Sleep(1 * time.Second)
		done <- true
	}()
	go func() {
		fmt.Println("Error GoROutine")
		time.Sleep(4 * time.Second)
		error <- true
	}()

	select {
	case <-done:
		fmt.Println("Done!!!!!")
	case <-error:
		fmt.Println("Error!!!!!")
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout!!!!!")

	}

	fmt.Println("Main function done...")

}
