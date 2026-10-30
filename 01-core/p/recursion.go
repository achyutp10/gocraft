package main

import "fmt"

func main() {
	fmt.Println(countdown(0))
	fmt.Println(countdown(1))
	countdown2(1, 4)

	fmt.Println(power(2, 3))
}

func countdown(n int) int {
	if n == 0 {
		return n
	}
	fmt.Println(n)
	return countdown(n - 1)
}

func countdown2(current, n int) {
	if current > n {
		return
	}

	fmt.Println(current)
	countdown2(current+1, n)

}

func power(m int, n int) int {
	if n == 0 {
		return 1
	}
	if m == 0 {
		return 0
	}
	return m * power(m, n-1)

}
