package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":  "#e4f5g6",
		"blue": "#e3g904",
	}
	delete(colors, "red")
	fmt.Println(colors)

}
