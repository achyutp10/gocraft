package main

import "fmt"

func main() {
	// var container [5]int
	// var no int
	// var sum int
	// var avg int
	// for i := 0; i < len(container); i++ {
	// 	fmt.Println("Enter no at index: ", i)
	// 	fmt.Scanln(&no)
	// 	container[i] = no
	// }
	// for i := 0; i < len(container); i++ {
	// 	// fmt.Println(container[i])
	// 	sum = sum + container[i]
	// 	avg = sum / len(container)
	// }
	// fmt.Println(sum)
	// fmt.Println(avg)

	// numbers := [6]int{34, 12, 89, 5, 67, 23}

	// largest := numbers[0]
	// smallest := numbers[0]

	// for i := 0; i < len(numbers); i++ {
	// 	if numbers[i] > largest {
	// 		largest = numbers[i]
	// 	}
	// 	if numbers[i] < smallest {
	// 		smallest = numbers[i]
	// 	}
	// }
	// fmt.Println("Smallest is: ", smallest)
	// fmt.Println("Largest is: ", largest)

	// for i := 0; i < len(numbers); i++ {
	// 	// fmt.Println(numbers[i])
	// 	sum = sum + numbers[i]
	// 	avg = sum / len(numbers)
	// }
	// fmt.Println(sum)
	// fmt.Println(avg)

	// var numContainer [3][3]int
	// var val int
	// for i := 0; i < len(numContainer); i++ {
	// 	for j := 0; j < len(numContainer); j++ {
	// 		fmt.Println("Enter no in position: ", i, " , ", j)
	// 		fmt.Scanln(&val)
	// 		numContainer[i][j] = val
	// 	}

	// }
	// for i := 0; i < len(numContainer); i++ {
	// 	fmt.Println()
	// 	for j := 0; j < len(numContainer); j++ {
	// 		fmt.Print(numContainer[i][j])
	// 	}
	// }

	arr := [5]int{1, 2, 3, 4, 5}
	var copiedArray *[5]int

	copiedArray = &arr
	copiedArray[0] = 100
	copiedArray[4] = 500

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}

	fmt.Println(arr)
	fmt.Println(*copiedArray)

}
