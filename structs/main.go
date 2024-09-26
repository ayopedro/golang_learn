package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"} // not recommended

	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// var alex person

	// alex.firstName = "Alexander"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Iyke",
		contactInfo: contactInfo{
			email: "jimiyke@test.com",
			zip:   1234,
		},
	}

	// jimPointer := &jim

	// jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")
	fmt.Printf("%+v", jim)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
