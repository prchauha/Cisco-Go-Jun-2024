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

// fmt.Stringer interface implementation
func (c Circle) String() string {
	return fmt.Sprintf("[Circle] Radius : %0.2f", c.Radius)
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

// fmt.Stringer interface implementation
func (r Rectangle) String() string {
	return fmt.Sprintf("[Rectangle] Length: %0.2f, Width: %0.2f", r.Length, r.Width)
}

// utility functions
/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area : ", val.Area())
	case Rectangle:
		fmt.Println("Area : ", val.Area())
	default:
		fmt.Println("argument does not support Area() method")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area : ", val.Area())
	default:
		fmt.Println("argument does not support Area() method")
	}
}
*/

/*
func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area : ", x.Area())
}
*/

// typed interface
type AreaFinder interface {
	Area() float64
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area : ", x.Area())
}

type PerimeterFinder interface {
	Perimeter() float64
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter : ", x.Perimeter())
}

// interface composition
type StatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintStats(x StatsFinder) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{Radius: 12}
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)

	r := Rectangle{Length: 10, Width: 12}
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)

	// PrintArea(100)

	// Using std lib interface
	fmt.Println(c)
	fmt.Println(r)
}
