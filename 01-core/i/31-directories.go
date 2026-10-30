package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// err := os.Mkdir("subdir", 0755)
	// checkError(err)
	checkError(os.Mkdir("subdir1", 0755))
	// defer os.RemoveAll("subdir1")

	os.WriteFile("subdir1/file", []byte(""), 0755)

	checkError(os.MkdirAll("subdir/parent/child", 0755))
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))

	os.WriteFile("subdir/parent/file", []byte(""), 0755)
	os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	checkError(os.Chdir("subdir/parent/child"))
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)
	result, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading subdir/parent/child")
	for _, entry := range result {
		fmt.Println(entry.Name())
	}

	checkError(os.Chdir("../../.."))
	dir, err = os.Getwd()
	checkError(err)
	fmt.Println(dir)

	// filepath.Walk and filepath.WalkDir
	pathFile := "subdir/parent/child"

	//
	fmt.Println("Walking Directory")

	err = filepath.WalkDir(pathFile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	checkError(err)

	// checkError(os.RemoveAll("subdir"))
	checkError(os.RemoveAll("./subdir"))
	// checkError(os.Remove("subdir1"))
	checkError(os.Remove("./subdir1"))

}
