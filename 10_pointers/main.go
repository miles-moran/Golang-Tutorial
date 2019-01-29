package main

import "fmt"

func main() {
	a := 5
	b := &a               //location of a rather than the value of a
	fmt.Printf("%T\n", b) //*int is  data type

	fmt.Println(a, b)

	fmt.Println(*b) //prints value at the location
	fmt.Println(*&a)

	*b = 10 //whatever is located at b (location) now has a value of 10
	fmt.Println(a)
}
