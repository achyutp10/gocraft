package main

import "fmt"

func main() {
	process(10)
}

func process(i int) {
	defer fmt.Println("Deffered i value: ", i)
	defer fmt.Println("First Deferred statement execued")
	defer fmt.Println("Second Deferred statement execued")
	defer fmt.Println("Third Deferred statement execued")
	i++

	fmt.Println("Noramal execution statement")
	fmt.Println(i)
}
