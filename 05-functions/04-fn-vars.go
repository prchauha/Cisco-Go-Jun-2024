package main

import "fmt"

func main() {
	// no parameters & no results
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi")
	}
	sayHi()
	sayHi = func() {
		fmt.Println("Hello!")
	}
	sayHi()

	// 1 parameter & no results
	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	// 2 parameter & no results
	var greetUser func(string, string)
	greetUser = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	greetUser("Magesh", "Kuppan")

	// 2 parameters & 1 result
	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	// 2 parameters & 2 results - using "named" results
	// "named" results are preferred when there are more than 1 return result (convention)
	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}
	fmt.Println(divide(100, 7))
	// receiving the results in variables
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}

// 2 parameters & 2 results
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/
