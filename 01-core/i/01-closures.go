package main

import "fmt"

func main() {
	// sequence := adder()
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	// sequence2 := adder()
	// fmt.Println(sequence2())

	subtractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	// Using the closure subtractor
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(2))
	fmt.Println(subtractor(3))
	fmt.Println(subtractor(4))
	fmt.Println(subtractor(5))

}

func adder() func() int {
	i := 0
	fmt.Println("Previous value if i: ", i)
	return func() int {
		i++
		fmt.Println("Adding 1 too i")
		return i
	}
}
