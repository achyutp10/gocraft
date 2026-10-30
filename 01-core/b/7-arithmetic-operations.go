package main

import "fmt"

func main() {
	a := 10
	b := 3
	var x, y int = 5, 2
	var result int

	const p float64 = 22 / 7.0
	fmt.Println("p =", p)

	result = x + y
	fmt.Println("x + y =", result)
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// overflow and underflow
	var maxInt int = 2147483647 // maximum value for a 32-bit signed integer
	fmt.Println("maxInt =", maxInt)
	maxInt++
	fmt.Println("maxInt after overflow =", maxInt)
	var minInt int = -2147483648 // minimum value for a 32-bit signed integer
	fmt.Println("minInt =", minInt)
	minInt--
	fmt.Println("minInt after underflow =", minInt)

	// overflow and underflow with unsigned integers
	var maxUint uint = 4294967295 // maximum value for a 32-bit unsigned integer
	fmt.Println("maxUint =", maxUint)
	maxUint++
	fmt.Println("maxUint after overflow =", maxUint)
	var minUint uint = 0 // minimum value for a 32-bit unsigned integer
	fmt.Println("minUint =", minUint)
	minUint--
	fmt.Println("minUint after underflow =", minUint)

}
