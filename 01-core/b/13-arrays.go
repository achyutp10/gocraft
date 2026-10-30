package main

import "fmt"

func main() {
	// var arrayName [size] elementType
	var numbers [5]int
	fmt.Println(numbers)

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array: ", fruits)
	fmt.Println("Third element: ", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray

	// copiedArray[0] = 100
	// fmt.Println("Original Array: ", originalArray)
	// fmt.Println("Copied Array: ", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index:", i, ":", numbers[i])
	}

	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// _ is blank identifier used to store unsued value
	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}
	a, _ := someFunction()
	fmt.Println(a)
	// fmt.Println(b)

	// b := 2
	// _ = b

	fmt.Println("The length of the numbers array is", len(numbers))

	// comparing arrays
	arra1 := [3]int{1, 2, 3}
	arra2 := [3]int{10, 2, 3}

	fmt.Println("Array1 is equal to Array2: ", arra1 == arra2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray

	copiedArray[0] = 100
	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Copied Array: ", *copiedArray)
}
func someFunction() (int, int) {
	return 1, 2
}
