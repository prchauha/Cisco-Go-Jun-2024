package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	defer func() {
		fmt.Println("	[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println(" [main] - recovered from panic")
		}
	}()
	fmt.Println("main started")
	doSomething()
	fmt.Println("main completed")
}

func doSomething() {
	defer func() {
		fmt.Println("	[doSomething] - deferred")
		if err := recover(); err != nil {
			fmt.Println(" [doSomething] - recovered from panic")
		}
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
