package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type app []string

func (a app) print() {
	for index, app := range a {
		fmt.Println(index, app)
	}
}

func (a app) toString() string {

	return strings.Join([]string(a), ",")
}
func getLength(length int) int {
	return length
}
func writeToFile(message string) {
	fmt.Println(message)
	ioutil.WriteFile("messageFile", []byte(message), 0666)
}
func ReadfromFile(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	new_strings := strings.Split(string(content), ",")
	length := len(new_strings)
	fmt.Println(new_strings)
	fmt.Println(length)
}
