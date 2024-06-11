package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divisor cannot be zero")

func main() {
	var divisor int
	for {
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		q, r := divide(100, divisor)
		fmt.Printf("Dviding 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	}
}

// 3rd party api (CANNOT change this code)
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient, remainder = x/y, x%y
	return
}
