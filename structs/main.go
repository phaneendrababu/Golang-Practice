package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 50010,
		},
	}

	jimPointer := &jim
	jimPointer.updateFirstName("jimmy")
	jim.print()

}

//No effect on firstName
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
