package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)

		// fmt.Println(<-ch)
		fmt.Println("3 sec Goroutine finished")
	}()

	go func() {
		ch <- 1
		time.Sleep(2 * time.Second)
		fmt.Println("2 sec Goroutine finished")
	}()
	go func() {
		// ch <- 1
		time.Sleep(3 * time.Second)
		fmt.Println("3 sec Goroutine finished")
	}()

	receiver := <-ch
	fmt.Println(receiver)
}
