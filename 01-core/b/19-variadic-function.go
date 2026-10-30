package main

import "fmt"

func main() {
	// ... Ellipsis
	// func funName(param1 type1, param2 type2, param3 ...type3) returnType {
	// 	fn body
	// }

	fmt.Println("Sum of 1,2,3: ", sum(1, 2, 3))

	statement, total := sum2("The sum of 1,2,3 is ", 1, 2, 3)
	fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 9}

	st, total3 := sum2("Sum is ", numbers...)
	fmt.Println(st, total3)

}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func sum2(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}
