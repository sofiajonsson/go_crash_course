package main

import "fmt"
// go can support anonymous functions that can form closures
// can define a func inline without having to name it

func adder() func(int) int {
	sum:= 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	for i := 0; i <10; i++ {
		//since we set it to adder we can pass in the i
		fmt.Println(sum(i))
	}


}
