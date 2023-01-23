package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://example.com/")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99998)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	fmt.Println(resp)
}
