package basics

import (
	"fmt"
	// foo "net/http"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	// resp, err := foo.Get("https://www.google.com")
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Print("HTTP response status: ", resp.Status)
}
