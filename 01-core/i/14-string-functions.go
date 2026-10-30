package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	// str := "Hello Go! Go"
	str := "Hello Go!"
	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[1:7])

	// String conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))

	// string splitting
	fruits := "apple, orange, banana"
	fruits1 := "apple, orange, banana"
	parts := strings.Split(fruits, ",")
	parts1 := strings.Split(fruits1, "-")
	fmt.Println(parts)
	fmt.Println(parts1)

	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "Go"))
	fmt.Println(strings.Replace(str, "Go", "hoo", 2))
	fmt.Println(strings.Replace(str, "Go", "hoo", 1))
	fmt.Println(strings.Contains(str, "Go?"))
	fmt.Println(strings.Contains(str, "G"))

	strwspace := `Hello Everyone!`
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))
	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	fmt.Println(strings.Repeat("Foo ", 4))
	fmt.Println(strings.Count("Foo Fooo Hello", "He"))
	fmt.Println(strings.HasPrefix("Foo Fooo Hello", "He"))
	fmt.Println(strings.HasSuffix("Foo Fooo Hello", "lo"))

	str5 := "Hel1lo, 123, Go! 11"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	str6 := "Hello, 世の中"
	fmt.Println(utf8.RuneCountInString(str6))

	// String builder
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	// Convert builder to a string
	rslt := builder.String()
	fmt.Println(rslt)

	// using write rune to add a character

	builder.WriteRune(' ')
	builder.WriteString("How are you")
	rslt = builder.String()
	fmt.Println(rslt)

	// Reset the builder
	builder.Reset()

	builder.WriteString("Starting fresh")
	rslt = builder.String()
	fmt.Println(rslt)

}
