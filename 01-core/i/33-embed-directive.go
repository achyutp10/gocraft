package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed basic

var basicFolder embed.FS

func main() {
	fmt.Println("Embeded content:", content)
	content, err := basicFolder.ReadFile("basic/2-hello.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Embeded file content:", string(content))

	err = fs.WalkDir(basicFolder, "basic", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
