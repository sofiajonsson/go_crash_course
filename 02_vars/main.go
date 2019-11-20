package main

import "fmt"

//global
// var name = "Sofia"
// name := "Sofia" ==> shorthand cannot be global
func main() {
	//using var

		var age = 26
		var isCool = true
		var size float32 = 2.3
	//Shorthand
		name := "Sofia"
		//comment out top sofia if want to assign several times as below 
		name, email := "Sofia", "sofia@me.com"

		fmt.Println(name, age, isCool, email)
		fmt.Printf("%T\n", size)
}
