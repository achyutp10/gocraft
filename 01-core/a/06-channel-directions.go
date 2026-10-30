package main

import "fmt"

func main() {
	ch := make(chan int)
	sendData(ch)
	receiveData(ch)

}

// Send only channel
func sendData(ch chan<- int) { // producer
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// Receive only channel
func receiveData(ch <-chan int) { //consumer
	for value := range ch {
		fmt.Println("Received:", value)
	}
}
