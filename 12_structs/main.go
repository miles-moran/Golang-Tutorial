package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (point receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

//TWO METHODS

//VALUE RECEIVERS -

//POINTER RECEIVERS -

func main() {
	person1 := Person{
		firstName: "Samantha",
		lastName:  "Smith",
		city:      "Boston",
		gender:    "f",
		age:       25,
	}

	person2 := Person{"Miles", "Moran", "St. Louis", "m", 26}

	fmt.Println(person1)
	fmt.Println(person2)

	//person1.age++

	fmt.Println(person1.firstName)

	person1.hasBirthday()
	person2.getMarried("Diaz")
	person1.getMarried("Steinberg")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
