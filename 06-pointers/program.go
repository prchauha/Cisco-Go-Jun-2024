package main

import "fmt"

func main() {
	var num int
	num = 100

	var numPtr *int
	numPtr = &num // address of "num"
	fmt.Printf("[main] %p => %d\n", numPtr, num)

	// access the value through the reference (aka pointer) - dereferencing
	x := *numPtr
	fmt.Println(x)

	fmt.Println("Before incrementing, num =", num)
	increment(&num)
	fmt.Println("After incrementing, num =", num)

}

func increment(val *int) /* no return result */ {
	fmt.Printf("[increment] &val = %p\n", val)
	/* increment the given value */
	*val += 1
}
