// MAIN TYPES
// string
// bool
// int
// int in8 int16 int32 64
// uint uin8 uint16 uint32 64
// byte - alias for uint8
// rune - alias for int32

package main

import "fmt"

func main() {
	var name string = "Miles"
	var name2 = "Miles" //inferred
	var age int32 = 26
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name, age)
	fmt.Printf("%T\n", age)

	const age2 = 25

	name3 := "Miles" // shorthand method cannot be used outside of a function
	name4, email := "Brad", "Brad@hotmail.com"

	fmt.Println(name3, name4, email)
}
