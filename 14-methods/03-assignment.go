/*
Create the following types
	Circle
		- Radius (float64)
	Rectangle
		- Length (float64)
		- Width (float64)

Create the following Methods for both the types
	- Area() float64
	- Perimeter() float64

Formula
	Circle Area = math.Pi * Radius * Radius
	Circle Perimeter = 2 * math.Pi * Radius

	Rectangle Area = Length * Width
	Rectangle Perimeter = 2 * (Length + Width)

In the main()
	Create instances of Circle & Rectangle and Print the Area & Perimiter

*/

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func main() {
	c := Circle{Radius: 12}
	fmt.Println("Area : ", c.Area())
	fmt.Println("Perimter : ", c.Perimeter())

	r := Rectangle{Length: 10, Width: 12}
	fmt.Println("Area : ", r.Area())
	fmt.Println("Perimter : ", r.Perimeter())
}
