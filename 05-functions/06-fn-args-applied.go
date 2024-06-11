/*
Modify the below program (without modifying the "add" & "subtract" functions) to diplay the following:

Operation Started
Add Result : 300
Operation Completed
Operation Started
Subtract Result : -100
Operation Completed
*/
package main

import "fmt"

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(i1, i2 int) {
		fmt.Println("Multiply Result :", i1*i2)
	}, 100, 200)
}

// refactor to avoid duplication in the log wrapper functions by applying "commonility variability technique"
func logOperation(operation func(int, int), x, y int) {
	fmt.Println("Operation Started")
	operation(x, y)
	fmt.Println("Operation Completed")
}

// wrapper functions with log capabilities
func logAdd(x, y int) {
	fmt.Println("Operation Started")
	add(x, y)
	fmt.Println("Operation Completed")
}

func logSubtract(x, y int) {
	fmt.Println("Operation Started")
	subtract(x, y)
	fmt.Println("Operation Completed")
}

// 3rd party library apis
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
