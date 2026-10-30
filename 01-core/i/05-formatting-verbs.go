package main

import "fmt"

func main() {
	// --- General Formatting Verbs
	// %v   Prints the value in the default format
	// %#v  Prints the value in Go-syntax format
	// %T   Prints the type of the value
	// %%   Prints the % sign

	i := 156_505.5
	str := "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
	fmt.Printf("%T\n", str)

	// --- Integer Formatting Verbs
	// %b   Base 2
	// %d   Base 10
	// %+d  Base 10 and always show sign
	// %o   Base 8
	// %O   Base 8, with leading 0o
	// %x   Base 16, lowercase
	// %X   Base 16, uppercase
	// %#x  Base 16, with leading 0x
	// %4d  Pad with spaces (width 4, right justified)
	// %-4d Pad with spaces (width 4, left justified)
	// %04d Pad with zeroes (width 4)

	// int := 18
	num := 42

	fmt.Printf("%b\n", num)
	fmt.Printf("%d\n", num)
	fmt.Printf("%+d\n", num)
	fmt.Printf("%o\n", num)
	fmt.Printf("%O\n", num)
	fmt.Printf("%x\n", num)
	fmt.Printf("%X\n", num)
	fmt.Printf("%#x\n", num)
	fmt.Printf("%4d|\n", num)
	fmt.Printf("%-4d|\n", num)
	fmt.Printf("%04d\n", num)

	// --- String Formatting Verbs
	// %s   Prints the value as plain string
	// %q   Prints the value as a double-quoted string
	// %8s  Width 8, right justified
	// %-8s Width 8, left justified
	// %x   Hex dump of byte values
	// % x  Hex dump with spaces

	word := "GoLang"

	fmt.Printf("%s\n", word)
	fmt.Printf("%q\n", word)
	fmt.Printf("%8s|\n", word)
	fmt.Printf("%-8s|\n", word)
	fmt.Printf("%x\n", word)
	fmt.Printf("% x\n", word)

	// --- Boolean Formatting Verbs
	// %t   Prints true or false

	isGoFun := true

	fmt.Printf("%t\n", isGoFun)
	fmt.Printf("%v\n", isGoFun)

	// --- Floating-Point Formatting Verbs
	// %f   Decimal notation
	// %.2f Two digits after the decimal point
	// %e   Scientific notation (lowercase e)
	// %E   Scientific notation (uppercase E)
	// %g   Compact representation
	// %G   Compact representation (uppercase exponent)
	// %b   Binary exponent notation
	// %10.2f Width 10, two decimal places
	// %010.2f Width 10, padded with zeroes

	price := 1234.5678

	fmt.Printf("%f\n", price)
	fmt.Printf("%.2f\n", price)
	fmt.Printf("%e\n", price)
	fmt.Printf("%E\n", price)
	fmt.Printf("%g\n", price)
	fmt.Printf("%G\n", price)
	fmt.Printf("%b\n", price)
	fmt.Printf("%10.2f|\n", price)
	fmt.Printf("%010.2f\n", price)

	// --- Character Formatting Verbs
	// %c   Prints the Unicode character
	// %q   Prints the quoted character
	// %U   Prints Unicode format (U+...)
	// %#U  Unicode format with character

	char := 'A'

	fmt.Printf("%c\n", char)
	fmt.Printf("%q\n", char)
	fmt.Printf("%U\n", char)
	fmt.Printf("%#U\n", char)

	// --- Pointer Formatting Verbs
	// %p   Prints the memory address in hexadecimal

	x := 100
	ptr := &x

	fmt.Printf("%p\n", ptr)
	fmt.Printf("%v\n", ptr)
	fmt.Printf("%T\n", ptr)

	// --- Slice and Array Formatting
	// %v   Prints elements in default format
	// %#v  Prints Go-syntax representation
	// %T   Prints the type

	nums := []int{10, 20, 30}
	arr := [3]int{1, 2, 3}

	fmt.Printf("%v\n", nums)
	fmt.Printf("%#v\n", nums)
	fmt.Printf("%T\n", nums)

	fmt.Printf("%v\n", arr)
	fmt.Printf("%#v\n", arr)
	fmt.Printf("%T\n", arr)

	// --- Struct Formatting
	// %v   Prints field values
	// %+v  Prints field names and values
	// %#v  Prints Go-syntax representation

	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Achyut", Age: 23}

	fmt.Printf("%v\n", person)
	fmt.Printf("%+v\n", person)
	fmt.Printf("%#v\n", person)

	// --- Error Formatting
	// %v   Prints the error message
	// %s   Prints the error message as a string
	// %q   Prints the error message quoted

	err := fmt.Errorf("something went wrong")

	fmt.Printf("%v\n", err)
	fmt.Printf("%s\n", err)
	fmt.Printf("%q\n", err)

	// --- Width and Precision
	// %8.2f  Width 8, two decimal places
	// %-8.2f Left justified, two decimal places
	// %.3s   First three bytes of a string

	fmt.Printf("%8.2f|\n", price)
	fmt.Printf("%-8.2f|\n", price)
	fmt.Printf("%.3s\n", word)
}
