package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error Opening file:", err)
		return
	}
	defer func() {
		fmt.Println("Closing open file")
		file.Close()
	}()

	fmt.Println("File opened successfully")

	// Read the contents of the opened file
	data := make([]byte, 1024) // Buffer to read data into
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error Reading file:", err)
		return
	}
	fmt.Println("File Content:", string(data))

	// -------------------------------------

	scanner := bufio.NewScanner(file)
	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error Reading file:", err)
		return
	}

}
