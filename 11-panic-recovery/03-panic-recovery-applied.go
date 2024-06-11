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
		if q, r, err := divideClient(100, divisor); err == nil {
			fmt.Printf("Dviding 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
			continue
		} else {
			fmt.Printf("Error occured: %v , please try again!\n", err)
		}
	}
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api (CANNOT change this code)
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient, remainder = x/y, x%y
	return
}
