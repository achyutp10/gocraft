package main

import "fmt"

func main() {
	// func fn_name(parameters list) () or returnType{
	// code block
	// return value
	// }

	// sum := add(1, 2)
	// fmt.Println(sum)
	// fmt.Println(add(1, 2))

	// func() {
	// 	fmt.Println("Hello Anonymous Function")
	// }()

	// greet := func() {
	// 	fmt.Println("Hello Anonymous Function")
	// }
	// greet()

	// operation := add
	// result := operation(3, 5)
	// fmt.Println(result)

	// passing a fn as an argument
	result := applyOperation(5, 3, add)
	fmt.Println("5 + 3 = ", result)

	// Returning and using a fn
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 x 2 = ", multiplyBy2(6))

}

func add(a int, b int) int {
	return a + b
}

// fn taking fn as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// fn returns a fn

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
