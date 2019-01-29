package main

import "fmt"

func main() {
	ids := []int{33, 2, 44, 11, 64, 2}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//_ when variable will not be used
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	emails := make(map[string]string)

	//assign
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

}
