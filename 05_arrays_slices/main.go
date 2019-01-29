package main

import "fmt"

func main() {
	var FruitArr [2]string

	FruitArr[0] = "Apple"
	FruitArr[1] = "Orange"

	fmt.Println(FruitArr)

	FruitArr2 := [2]string{"Apple", "Orange"}

	fmt.Println(FruitArr2)

	FruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(FruitSlice)
	fmt.Println(len(FruitSlice))
	fmt.Println(FruitSlice[1:3])

}
