/* share memory by communicating */
package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("Result :", result)
}

// producer
func add(x, y int, ch chan int) {
	time.Sleep(4 * time.Second)
	result := x + y
	ch <- result
}
