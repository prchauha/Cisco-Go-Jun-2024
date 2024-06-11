package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	divisor := 6
	// divisor := 0

	if q, r, err := divide(100, divisor); err == ErrDivideByZero {
		fmt.Println("Do not attempt to divide by 0")
	} else if err != nil {
		fmt.Printf("Error :%v\n", err)
	} else {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	}
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be zero")
		return 0, 0, err
	}
	quotient, remainder := x/y, x%y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		// err = errors.New("divisor cannot be zero")
		err = ErrDivideByZero
		return
	}
	if rand.Intn(20)%2 == 0 {
		err = errors.New("some dummy error")
		return
	}
	quotient, remainder = x/y, x%y
	return
}
