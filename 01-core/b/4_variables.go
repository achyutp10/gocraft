package main

import "fmt"

// middleName := "William" // short variable declaration this will cause an error because short variable declarations can only be used inside functions
var middleName string = "William" // initialized variable this is a package-level variable and can be accessed from any function within the package

func main() {
	middleName := "ailliam" // short variable declaration

	fmt.Println("Variables:")
	var x int = 5 // initialized variable
	fmt.Println("x:", x)
	x = 10 // reassigning variable
	fmt.Println("x:", x)
	var name string = "John" // initialized variable
	fmt.Println("name:", name)
	var age int // uninitialized variable
	age = 30    // assigning value to uninitialized variable
	fmt.Println("age:", age)
	fmt.Println(true && true)

	count := 0 // short variable declaration
	fmt.Println("count:", count)
	count++
	fmt.Println("count:", count)
	lastName := "Doe" // short variable declaration
	fmt.Println("lastName:", lastName)

	// Default values
	// Numeric types default to 0, booleans default to false, and strings default to an empty string.
	var defaultInt int
	var defaultBool bool
	var defaultString string
	fmt.Println("defaultInt:", defaultInt)
	fmt.Println("defaultBool:", defaultBool)
	fmt.Println("defaultString:", defaultString)

	// Pointers, slices, maps, channels, and function types default to nil.
	var defaultPointer *int
	var defaultSlice []int
	var defaultMap map[string]int
	var defaultChannel chan int
	var defaultFunction func()
	fmt.Println("defaultPointer:", defaultPointer)
	fmt.Println("defaultSlice:", defaultSlice)
	fmt.Println("defaultMap:", defaultMap)
	fmt.Println("defaultChannel:", defaultChannel)
	fmt.Println("defaultFunction:", defaultFunction)

	// Scope
	// Variables declared inside a function are local to that function and cannot be accessed outside of it.
	// fmt.Println(firstName) // This will cause an error because firstName is not defined in this scope
	fmt.Println(middleName) // This will work because middleName is defined in this scope

}

func printName() {
	middleName := "mayor"

	firstName := "John" // local variable
	fmt.Println("firstName:", firstName)
	fmt.Println("middleName:", middleName) // This will work because middleName is a package-level variable
}
