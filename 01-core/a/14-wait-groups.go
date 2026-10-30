package main

import (
	"fmt"
	"sync"
	"time"
)

// construction example

type Worker struct {
	ID   int
	Task string
}

// Perform Task simulates a worker performing a task
func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("WorkerID %d started %d\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func main() {
	var wg sync.WaitGroup
	// Define tasks to be performed by workers
	tasks := []string{"diging", "Laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i + 1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Construction is fnished
	fmt.Println("Construction finished")

}

// example with channels

func worker1(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("WorkerID %d starting.\n", id)
	time.Sleep(time.Second) // simulate work
	for task := range tasks {
		results <- task * 2
	}

	results <- id * 2
	fmt.Printf("Worker id %d finished\n", id)
}

func main1() {
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 5
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker1(i+1, tasks, results, &wg)
	}

	for i := range numJobs {
		tasks <- i + 1
	}

	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Results:", result)
	}
}

// --- basic example without channels

// func worker2(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) // simulate processing
// 	fmt.Printf("Worker %d finished\n", id)
// }
// func main2() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	wg.Add(numWorkers) // correct way of adding counter to wait groups

// 	// launch workers

// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait() // blocking mechanism

// 	fmt.Println("All Workers finished")
// }
