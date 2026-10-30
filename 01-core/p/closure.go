package main

import "fmt"

func main() {

	double := createMultiplier(2)
	triple := createMultiplier(3)
	fmt.Println(double(5))
	fmt.Println(triple(5))

	greet := createGreeter("Achyut")
	fmt.Println(greet())
	fmt.Println(greet())

	next := createSequence(10, 5)
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

}

func createMultiplier(val int) func(int) int {
	return func(x int) int {
		val = val * x
		return val
	}
}

func createGreeter(greet string) func() string {
	return func() string {
		return "Hello " + greet
	}
}

func createSequence(x int, y int) func() int {
	return func() int {
		i := x
		x = x + y
		return i
	}

}
