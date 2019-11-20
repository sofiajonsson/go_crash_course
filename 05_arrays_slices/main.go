package main

import "fmt"

func main() {
	//Arrays : need to name types and they are fixed length so that's where slice comes in
		// var fruitArray [2]string

	//Assign values
		// fruitArray[0]= "Apple"
		// fruitArray[1]= "Orange"

	// Declare and assign
		// fruitArray := [2]string{"Apple", "Orange"}
		//
		// fmt.Println(fruitArray)
		// fmt.Println(fruitArray[1])

	// Slices: don't include index or it will throw index "num" is out of bounds
		fruitSlice:= []string{"Apple", "Orange", "Grape"}

		fmt.Println(fruitSlice)
		fmt.Println(len(fruitSlice))
		//prints within the range
		fmt.Println(fruitSlice[1:2])
}
