package main

import "fmt"

func main() {

	colors := map[string]string{
		"green": "#e0e0e0",
	}

	colors["red"] = "#e4f5g6"
	colors["blue"] = "#e3g904"
	colors["white"] = "#fffff"

	//fmt.Println(colors)
	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hexCode := range colors {
		fmt.Println("hexcode of ", color, "is :", hexCode)
	}
}
