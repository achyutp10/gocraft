package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He asid, \"I am great\"")
	fmt.Println(`He asid, "I am great"`)

	// Compile a regex pattern to match email

	re := regexp.MustCompile(`[a-zA-Z0-9,_+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// test strings
	email1 := "user@email.com"
	email2 := "invalid_email"

	// Match
	fmt.Println("Email1:", re.MatchString(email1))
	fmt.Println("Email2:", re.MatchString(email2))

	// capturing groups
	// Compile a regex pattern to capture date components

	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// Test strong
	date := "2024-07-30"

	// Find all submatches

	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// Source string

	str := "Hello World"
	re = regexp.MustCompile(`[aeiou]`)
	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// i -case insensitive
	// m - multline model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)
	// Test string
	text := "Golang is great"

	// Match
	fmt.Println("Match:", re.MatchString(text))
}
