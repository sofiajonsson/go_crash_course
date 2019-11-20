package main

import "fmt"
// pointer lets you point to a value in memory where this value is stored
// Reason you would use pointers is bc you might have to pass a lot of data stored at an address and if you pass address instead of data
//it can increase performance and to just change values at specific locations
// everything in go is passed by value
func main() {
		a := 5
		//ampersand is pointer of a and prints out memory address not value
		b := &a

		fmt.Println(a, b)
		fmt.Printf("%T\n", a) // returns int
		fmt.Printf("%T\n", b) // returns * int

		// Use * to read val from address
		// Can also run (*&b)
		fmt.Println(*b)

		// Change val with pointer
		*b = 10
		fmt.Println(a)


}
