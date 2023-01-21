package main

import "fmt"

func main() {
	fmt.Println("Hello")

	apps := app{"Facebook", "Instagram", "WhatsApp"}

	apps.print()
	message := apps.toString()
	writeToFile(message)
	ReadfromFile("messageFile")
	/*for index, app := range apps {
		fmt.Println(index, app)
	}*/
}
