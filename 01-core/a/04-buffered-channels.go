package main

import (
	"fmt"
	"time"
)

func main1() {
	// ======== Blocking on receive only if the buffer is empty
	ch := make(chan int, 2)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		ch <- 1
	}()
	fmt.Println("Value:", <-ch)
	fmt.Println("Value:", <-ch)
	fmt.Println("End of program")
}

func main() {

	// =================== Blocking on send only if the buffer is full
	// make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	println("Receiving from buffer")
	fmt.Println("Value", <-ch)
	fmt.Println("Value", <-ch)

	go func() {
		println("Goroutine 2 sec timer started")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch) // ends <- starts
	}()
	fmt.Println("Blocking Starts")
	ch <- 3 // blocks because buffer is full
	fmt.Println("Blocking Ends")
	fmt.Println("Received:", <-ch)
	fmt.Println("Value", <-ch)
	fmt.Println("Buffered Channel")
}
