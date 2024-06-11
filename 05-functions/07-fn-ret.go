package main

import (
	"fmt"
	"math/rand"
)

// executor - executes the given function
func main() {
	fn := getFn()
	fn()
}

// decider that decides "which" function to execute
func getFn() func() {
	randomNo := rand.Intn(100)
	fmt.Println(randomNo)
	switch {
	case randomNo%2 == 0:
		return f1
	case randomNo%3 == 0:
		return f2
	default:
		return func() {
			fmt.Println("anon fn invoked")
		}
	}
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
