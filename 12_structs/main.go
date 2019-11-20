package main

import (
	"fmt"
	"strconv"
)

// Very important part of go
// Basically like classes, no classes in go
// create a struct for a person and then create props like name, age, email
// create methods associated with that struct (2 main types)
	// pointer recievers
	//value recievers


// Define person struct
type Person struct {
	// //define properties of struct
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int
		// to shorten datatypes:
		firstName, lastName, city, gender string
		age int
	}



// Greeting method (value reciever) (we arent changing anything)
// p is specified identifier, Person is the name of the struct. p is similar to "this"
//will receive mis-matched data types error if concatenating with strings and int so import new package to use strconv
 	func (p Person) greet() string {
		return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
	}

// hasBirthday method (pointer reciever) ==>> will actually change something and use "*"
  	func (p *Person) hasBirthday(){
			p.age++
		}

// getMarried (pointer reciever)
	func (p *Person) getMarried(spouseLastName string){
		if p.gender == "m" {
			return
		} else {
			p.lastName = spouseLastName
		}
	}

func main() {
		// Init person using struct
		 person1 := Person{firstName: "Samantha", lastName: "Jones", city: "NYC", gender: "f", age: 39}

		 //Alternative
		  person2 := Person{"Mister", "Smithers", "NYC", "m", 79}

		 // fmt.Println(person1.firstName)
		 // //can increment property and have it reflected in bottom fmt
		 // person1.age++
		 // fmt.Println(person1)

		  person1.hasBirthday()
			person1.getMarried("Williams")
			person2.getMarried("Williams")
		 // person1.hasBirthday()
		 // person1.hasBirthday()
		 //Samantha is now 42 bc her bday has incremented so many time
		 fmt.Println(person1.greet())
		 fmt.Println(person2.greet())


//Methods: value recievers(for methods doing calculations and stuff, dont actually change anything ) and pointer recievers (actually change things)
	// value recievers (don't go inside of struct )
}
