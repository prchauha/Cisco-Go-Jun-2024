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

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap( /* ? */ )
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(val *int) /* no return result */ {
	fmt.Printf("[increment] &val = %p\n", val)
	/* increment the given value */
	*val += 1
}

func swap() /* no return reuslts */ {
	/* swap the given 2 values */
}
