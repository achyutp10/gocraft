package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numInt)
	fmt.Println(numInt + 1)

	numiInt, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numiInt)
	fmt.Println(numiInt + 1)

	floatStr := "3.14"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(floatVal)
	fmt.Println(floatVal + 0.1)

	binaryStr := "1010"
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decimal)
	fmt.Println(decimal + 11)

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hex)
	fmt.Println(hex + 12)

	invalidNum := "456abc"
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(invalidParse)
	fmt.Println(invalidParse + 12)

}
