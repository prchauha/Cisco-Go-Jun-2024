package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //schedule the f1() execution through the scheduler and the execution MAY happen in future
	f2()

	// blocking the execution so that the scheduler can look for other goroutines scheduled and execute them

	// DO NOT DO THIS IN PRODUCTION
	time.Sleep(2 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
