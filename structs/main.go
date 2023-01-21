package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contcat   contactInfo
}

/*type person struct {
	firstName string
	lastName  string
}*/

func main() {
	/*alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)*/
	/*var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)*/

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contcat: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 50010,
		},
	}

	fmt.Println(jim)
	fmt.Printf("%+v", jim)
}
