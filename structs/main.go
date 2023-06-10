package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	relax := person{lastName: "Sahu", firstName: "Saroj"}
	fmt.Println(relax)

	newPerson := person{}
	newPerson.firstName = "someone"
	newPerson.lastName = "else"
	fmt.Println(newPerson)
	fmt.Printf("%+v", newPerson)
	fmt.Println()
	newPerson.contact = contactInfo{
		email: "abc@xyz.com",
		zip:   12345,
	}

	fmt.Println(newPerson)

	relax.print()
	relax.updateFirstName("Yahoo")
	relax.print()

	oenone := person{
		firstName: "Onemore",
		lastName:  "Guys",
		contact: contactInfo{
			email: "wow@xyz.com",
			zip:   54321,
		},
	}
	oenone.print()

	pointerToOneMorePerson := &oenone
	pointerToOneMorePerson.print()
	pointerToOneMorePerson.updateFirstNames("Hello again")
	pointerToOneMorePerson.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p *person) updateFirstName(firstName string) {
	p.firstName = firstName
}

func (pointerToUpdate *person) updateFirstNames(firstName string) {
	pointerToUpdate.firstName = firstName
}
