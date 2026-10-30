package main

import (
	"fmt"
	"time"
)

// Basic worker pool pattern
func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		// Simulate work

		time.Sleep(time.Second)
		results <- task * 2
	}
}

func main1() {
	numWorkers := 3
	numJobs := 10
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// create workers

	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	// Sent values to the tasks channel
	for i := range numJobs {
		tasks <- i
	}
	close(tasks)

	// collect the results
	for range numJobs {
		result := <-results
		fmt.Println("Result:", result)
	}

}

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate processing ticket request
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %d", req.numTickets, req.personID, req.cost)
		// Simulate processing time
		time.Sleep(time.Second)
		results <- req.personID

	}
}

func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	// Start ticket processor/worker
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send ticket requests
	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for persionID %d processed successfully!\n", <-ticketResults)
	}

}
