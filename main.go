package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contact{
			email:   "jimparty@gmail.com",
			zipCode: 89423,
		},
	}
	jim.changeName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) changeName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
