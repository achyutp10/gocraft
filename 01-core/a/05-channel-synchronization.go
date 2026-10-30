package main

import (
	"fmt"
	"time"
)

func main1() {
	done := make(chan struct{})

	go func() {
		fmt.Println("Working")
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	<-done
	fmt.Println("Finished")
}

func main2() {
	ch := make(chan int)

	go func() {
		ch <- 9 // bloking until the value is received
		fmt.Println("Sent value")
	}()
	value := <-ch
	fmt.Println(value) // bloking until the value is sent
}

// SYNCHRONIZING MULTIPLE GOROUTINES AND ENSURING THAT ALL GOROUTINES ARE COMPLETE
func main3() {
	numGoRoutines := 3
	done := make(chan int, 3)

	for i := range numGoRoutines {
		time.Sleep(2 * time.Second)
		go func(id int) {
			fmt.Printf("Goroutine %d working...\n", id)
			done <- 1 // sending signal of completion
		}(i)
	}

	for range numGoRoutines {
		<-done // Wait for each goroutine to finish, WAIT for all goroutine to signal completion
	}

	fmt.Println("All goroutine are complete")
}

// SYNCHRONIZING data exchange

func main() {
	data := make(chan string)
	go func() {
		for i := range 5 {
			data <- "hello " + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()
	// close(data) // channel closed before goroutine could send a value to the channel

	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	} // loops over only on active channel, create receiver each time and stops creating receiver (looping) once
	// the channel is closed
}
