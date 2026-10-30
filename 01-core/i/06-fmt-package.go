package main

import "fmt"

func main() {

	// --- Printing Functions

	// Print: prints without automatically adding a newline
	fmt.Print("Hello")
	fmt.Print("World")
	fmt.Print(12, 456)
	fmt.Println()

	// Println: prints arguments separated by spaces and adds a newline
	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println(12, 456)

	// Printf: prints using format specifiers
	name := "John"
	age := 25

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// --- Formatting Functions

	// Sprint: returns a formatted string without adding a newline
	s := fmt.Sprint("Hello ", "World!", 123, 456)
	fmt.Print(s)
	fmt.Println()

	// Sprintln: returns a string with spaces between arguments
	// and a newline at the end
	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Print(s)
	fmt.Print(s)

	// Sprintf: returns a formatted string using format specifiers
	sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(sf)
	fmt.Println(sf)

	// --- Scanning Functions

	// Scan: reads whitespace-separated values from standard input
	var firstName string
	var inputAge int

	fmt.Print("Enter your name and age: ")
	fmt.Scan(&firstName, &inputAge)
	fmt.Printf("Name: %s, Age: %d\n", firstName, inputAge)

	// Scanln: reads values until a newline
	// It stops scanning at a newline
	var city string
	var country string

	fmt.Print("Enter city and country: ")
	fmt.Scanln(&city, &country)
	fmt.Printf("City: %s, Country: %s\n", city, country)

	// Scanf: reads input according to a format string
	var num1 int
	var num2 int

	fmt.Print("Enter two numbers separated by a comma (e.g., 10,20): ")
	fmt.Scanf("%d,%d", &num1, &num2)
	fmt.Printf("First: %d, Second: %d\n", num1, num2)

	// --- Error Handling with Scanning Functions

	var number int

	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&number)

	if err != nil {
		fmt.Println("Invalid input:", err)
	} else {
		fmt.Println("You entered:", number)
	}

	// --- String Formatting and Scanning

	// Sscan: scans whitespace-separated values from a string
	var username string
	var userAge int

	input := "Alice 30"
	fmt.Sscan(input, &username, &userAge)
	fmt.Printf("Username: %s, Age: %d\n", username, userAge)

	// Sscanf: scans values from a string using a format
	var product string
	var price float64

	productInput := "Laptop 85000.50"
	fmt.Sscanf(productInput, "%s %f", &product, &price)
	fmt.Printf("Product: %s, Price: %.2f\n", product, price)

	// Sscanln: scans values from a string until a newline
	var a string
	var b string

	line := "Hello World"
	fmt.Sscanln(line, &a, &b)
	fmt.Printf("A: %s, B: %s\n", a, b)

	// --- Writing to Other Destinations

	// Fprint: writes to a destination such as a file or standard output
	// fmt.Fprint(fmt.Stdout, "Written using Fprint\n")

	// Fprintln: writes arguments separated by spaces and adds a newline
	// fmt.Fprintln(fmt.Stdout, "Hello", "from", "Fprintln")

	// Fprintf: writes formatted output to a destination
	// fmt.Fprintf(fmt.Stdout, "Name: %s, Age: %d\n", name, age)

	// Error Formatting Function
	errrr := checkAge(15)
	if err != nil {
		fmt.Println("Error: ", errrr)
	} else {
		fmt.Println("You are elligible")
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive", age)
	}
	return nil
}
