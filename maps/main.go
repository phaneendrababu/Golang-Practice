package main

import "fmt"

func main() {

	colors := make(map[string]string)

	colors["red"] = "#e4f5g6"
	colors["blue"] = "#e3g904"
	colors["white"] = "#fffff"

	//fmt.Println(colors)
	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hexCode := range colors {
		fmt.Println("hexcode of ", color, "is ", hexCode)
	}
}
