package main

import (
	"fmt"
	"time"
)

// Goroutines are just fn that leave the main thread and run in background and come back
// to join the main thread once the
// fns are finished/ready to retuurn value dont stop program flow and are non blocking like async await like js

func main() {
	var err error

	fmt.Println("Begin")
	go sayHello()
	fmt.Println("After say hello")
	// err = doWork()
	// err = go doWork() // error not accepted
	go func() {
		err = doWork()
	}()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed")
	}
	time.Sleep(4 * time.Second)

}

func sayHello() {
	time.Sleep(time.Second * 1)
	fmt.Println("Hello from Goroutine")
	time.Sleep(time.Second * 2)

	go printNos()
	go printLetters()
	fmt.Println("Hello from Goroutine after 2 sec")
}

func printNos() {
	for i := 0; i < 5; i++ {
		fmt.Println("Numbers:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("I am an error")
}
