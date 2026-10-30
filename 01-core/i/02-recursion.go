package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
	fmt.Println(sumOfDigits(9))
	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigits(12345))

}

func factorial(n int) int {
	// Base case: facotorial of 0 is 1

	if n == 0 {
		return 1
	}
	// Recursive case: factorla of n is n * factorial(n-1)
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	// Base case:
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}
