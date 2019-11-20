package main

import "fmt"

func main() {
	// Define map with [key]value output
		// emails := make(map[string]string)

	// Assign key
		// emails["Bob"] = "bob@gmail.com"
		// emails["Sharon"] = "sharon@gmail.com"
		// emails["Mike"] = "mike@gmail.com"

	//Declare map and add key
			emails := map[string]string{"Bob": "bob@gmail.com", "Sharon" : "sharon@gmail.com"}
			emails["Mike"] = "mike@gmail.com"
	fmt.Println(emails)
	//check length
	fmt.Println(len(emails))
	//would most likely be searching by key since there would be many Bobs in db
	fmt.Println(emails["Bob"])


	//Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

}
