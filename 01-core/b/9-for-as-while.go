package main

import "fmt"

func main() {
	// i := 1
	// for i <= 5 {
	// 	fmt.Println("Iteration:", i)
	// 	i++
	// }

	// for {
	// 	fmt.Println("Hello")
	// }

	// sum := 0

	// For as while
	// for {
	// 	sum += 10
	// 	fmt.Print("Sum:", sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Numberr: ", num)
		num++
	}
}
