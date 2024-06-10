package main

import "fmt"

func main() {

	// no parameters & no results
	func() {
		fmt.Println("Hi")
	}()

	// 1 parameter & no results
	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	// 2 parameter & no results
	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	//TODO: replace the following with anonymous functions

	/*
		fmt.Println(func(x, y int) int {
			return x + y
		}(100, 200))
	*/

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(result)

	// fmt.Println(divide(100, 7))
	// receiving the results in variables
	q, r := func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}(100, 7)

	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	// receiving partial results

	q, _ = func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)

}
