package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {

	for range 20 {
		// check for an error
		if err := doSomething(); err != nil {
			fmt.Println("Error :", err)
		} else {
			fmt.Println("All good!")
		}
	}

}

func doSomething() error {
	// error variable
	var err error

	// creating an error
	err = errors.New("unintentional error")

	if rand.Intn(100)%2 == 0 {
		// returing the error
		return err
	}

	return nil
}
