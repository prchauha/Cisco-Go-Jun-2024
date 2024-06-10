package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	greetUser("Magesh", "Kuppan")
	fmt.Println(add(100, 200))

	// fmt.Println(divide(100, 7))

	// receiving the results in variables

	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	// receiving partial results
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
	*/
}

// no parameters & no results
func sayHi() {
	fmt.Println("Hi")
}

// 1 parameter & no results
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

// 2 parameter & no results
/*
func greetUser(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greetUser(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

// 2 parameters & 1 result
func add(x, y int) int {
	return x + y
}

// 2 parameters & 2 results
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// 2 parameters & 2 results - using "named" results
// "named" results are preferred when there are more than 1 return result (convention)
func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
