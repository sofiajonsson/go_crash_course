package main

import "fmt"

// imput type and return type
func greeting(name string) string {
	return "Hello " + name
}

//dont need to specify int for both since theyre the same
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
		fmt.Println(getSum(3, 4))
}
