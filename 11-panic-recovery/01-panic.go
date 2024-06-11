package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	defer func() {
		fmt.Println("	[main] - deferred")
	}()
	fmt.Println("main started")
	doSomething()
	fmt.Println("main completed")
}

func doSomething() {
	defer func() {
		fmt.Println("	[doSomething] - deferred")
	}()
	fmt.Println("started doing something")
	unpredictableFunc()
	fmt.Println("completed doing something")
}
func unpredictableFunc() {
	defer func() {
		fmt.Println("	[unpredictableFunc] - deferred")
	}()
	if rand.Intn(100)%2 == 0 {
		panic(errors.New("random error"))
	}
}
