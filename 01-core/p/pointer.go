package main

import "fmt"

func main() {

	var x *int
	var a int = 20

	x = &a
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(*x)

	x := 20
	var a *int
	a = &x
	*a = 50
	fmt.Println(x)

	x, y := 10, 20
	swap(&x, &y)
	fmt.Println(x, y)

	double(&x)
	fmt.Println(x)

	//
	arr := [4]int{2, 4, 6, 8}
	doubleArray(&arr)
	fmt.Println(arr) // Output: [4 8 12 16]

}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func double(p *int) {
	*p = *p * 2
}

func doubleArray(p *[4]int) {
	for i := 0; i < len(p); i++ {
		p[i] = p[i] * 2
	}
}
