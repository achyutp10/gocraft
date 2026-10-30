package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!"
	rawMessage := `Hello Go`
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message variable is", len(message))
	fmt.Println("Length of message2 variable is", len(message2))
	fmt.Println(message[0]) // ASCII value

	greeting := "Hello "
	name := "Alice"
	msg := greeting + name
	fmt.Println(msg)
	fmt.Println(greeting + name)

	str := "apple"
	str1 := "Apple"
	str2 := "Banana"
	str3 := "app"
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str2)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%x\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))
	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	// runes

	var ch rune = 'a'
	jch := 'カ'
	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T", ch)
	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "日本語" // japanese text
	fmt.Println(NIHONGO)

	jhello := "日本語" // japanese text
	for _, runeVal := range jhello {
		fmt.Printf("%c\n", runeVal)
		fmt.Printf("%v\n", runeVal)
	}

	r := '😊'
	fmt.Println("%v\n", r)
	fmt.Println("%c\n", r)

}
