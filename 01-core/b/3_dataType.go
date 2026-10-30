package basics

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Integers:", 42)
	fmt.Println("Floating Point Numbers:", 3.14)
	fmt.Println("Booleans:", true)
	fmt.Println("Strings:", "Hello, World!")
	fmt.Println("Complex Numbers:", 3+4i)
	fmt.Println("Constants:", 3.14, "Pi")
	fmt.Println("Arrays:", [3]int{1, 2, 3})
	fmt.Println("Structs:", struct{ Name string }{Name: "John"})
	fmt.Println("Pointers:", new(int))
	fmt.Println("Maps:", map[string]int{"John": 25})
	fmt.Println("Functions:", func() {})
	fmt.Println("Channels:", make(chan int))
	fmt.Println("JSON:", `{"name": "John", "age": 30}`)
	fmt.Println("Text and HTML Elements:", "<h1>Hello, World!</h1>")

}
