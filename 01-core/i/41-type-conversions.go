package main

import "fmt"

func main() {
	var a int = 32
	b := int32(a)
	c := float64(b)
	// d := bool(true)
	// d := bool("correct")

	e := 3.14
	f := int(e)
	fmt.Println(f, c)

	// Type(value)

	g := "Hello"
	var h []byte
	h = []byte(g)
	fmt.Println(h)

	// i := []byte{255, 72} // take uint8 upto 255 only
	i := []byte{'a', 'b'}
	j := string(i)
	fmt.Println(j)
}
