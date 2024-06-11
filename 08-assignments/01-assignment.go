/*
Refactor the below into functions
- Ensure that each function has only 1 reason to change (Single Responsibility Principle)
- Use higher order functions wherever necessary
*/

package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

var operations map[int]func(int, int) int = make(map[int]func(int, int) int)

func main() {
	var userChoice int
	prepareOperationCatalogue()
	for {
		userChoice = getUserChoice()
		if userChoice == 0 {
			break
		}
		processUserChoice(userChoice)
	}
}

func prepareOperationCatalogue() {
	addOperation(1, add)
	addOperation(2, subtract)
	addOperation(3, multiply)
	addOperation(4, divide)
	addOperation(5, power)
}

func addOperation(operationId int, operation func(int, int) int) {
	operations[operationId] = operation
}

func processUserChoice(userChoice int) {
	operation := getOperation(userChoice)
	if operation == nil {
		fmt.Println("Invalid choice")
		return
	}
	n1, n2 := getOperands()
	result := operation(n1, n2)
	fmt.Println("Result :", result)
}

func getOperation(userChoice int) func(int, int) int {
	return operations[userChoice]
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func getUserChoice() int {
	var userChoice int
	for id, operation := range operations {
		fmt.Printf("%d. %s\n", id, getFunctionName(operation))
	}
	fmt.Println("0. Exit")
	fmt.Println("Enter your choice :")
	fmt.Scanln(&userChoice)
	return userChoice
}

func getOperands() (n1, n2 int) {
	fmt.Println("Enter the operands:")
	fmt.Scanln(&n1, &n2)
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func power(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
