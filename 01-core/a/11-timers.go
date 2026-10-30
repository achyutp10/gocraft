package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(1 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("Timer1 expired")
	case <-timer2.C:
		fmt.Println("Timer2 expired")
	}
}

// Schedule delayed operations

func main3() {
	timer := time.NewTimer(2 * time.Second) // non blocking timer starts

	go func() {
		<-timer.C
		fmt.Println("Delayed operation executed")
	}()

	fmt.Println("Waiting...")

	time.Sleep(3 * time.Second) // blocking timer starts
	fmt.Println("End of the program")
}

// Timeout
func longRunningOperations() {
	for i := range 20 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main2() {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)

	go func() {
		longRunningOperations()
		done <- true
	}()

	select {
	case <-timeout:
		fmt.Println("Operation timed out")

	case <-done:
		fmt.Println("Operation completed")

	}
}

// Basic timer user
func main1() {
	time.Sleep(time.Second)
	fmt.Println("Starting app")
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting for timer .C")

	stopped := timer.Stop()
	if stopped {
		fmt.Println("Timer stopped")
	}

	timer.Reset(time.Second)
	fmt.Println("Timer reset")
	<-timer.C // blocking in nature
	fmt.Println("Timer expired")
}
