package main

import "fmt"

// executor - executes the given function
func main() {
	exec(f1) //=> execute f1
	exec(f2) //=> execute f2
	exec(func() {
		fmt.Println("anon fn invoked")
	})
}

// decider that decides "which" function to execute
func get(fnRef func()) {
	fnRef()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
