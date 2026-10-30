package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0, 5)
	var num int
	for i := 0; i < cap(slice); i++ {
		fmt.Printf("Enter at position %d: ", i)
		fmt.Scanln(&num)
		slice = append(slice, num)
	}
	fmt.Println("Length: ", len(slice))
	fmt.Println("Capacity: ", cap(slice))
	slice = append(slice, 60, 70)
	fmt.Println("Updated slice: ", slice)
	fmt.Println("Updated Length: ", len(slice))
	fmt.Println("Updated Capacity: ", cap(slice))
	slice[2] = 99
	fmt.Println("Final slice: ", slice)

	sub := slice[1:4]
	fmt.Println(sub)

}
