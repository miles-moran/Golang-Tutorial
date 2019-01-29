package main

import "fmt"

func main() {

	//define map
	emails := make(map[string]string)

	//assign
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails)

	emails2 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(emails2)
}
