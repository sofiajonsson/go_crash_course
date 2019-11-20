package main

import (
	"fmt"
	"math"
)

// datatypes that represent a set of method signatures for structs
// you can define an interface ex: shape that has a method named area
//and then we can then use that interface on a struct called rectangle or circle

//Define interface : shape with area method that returns a float 64
type Shape interface {
	area() float64
}

//pass in props x, y, radius
type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, height float64
}

// create area method for both of these structs (circle and rectangle)
	// value reciever so no * needed
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Will return area
func getArea(s Shape) float64 {
	return s.area()
}


func main() {
	circle := Circle{x:0, y:0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))


}
