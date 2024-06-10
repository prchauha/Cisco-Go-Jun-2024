package main

import "fmt"

func main() {
	/*
		const Red int = 0
		const Green int = 1
		const Blue int = 2
	*/

	/*
		const (
			Red int = 0
			Green int = 1
			Blue int = 2
		)
	*/

	/*
		const (
			Red   int = iota
			Green int = iota
			Blue  int = iota
		)
	*/

	/*
		const (
			Red   = iota
			Green = iota
			Blue  = iota
		)
	*/

	/*
		const (
			Red   = iota // 0
			Green        // 1
			Blue         // 2
		)
	*/

	/*
		const (
			Red   = iota + 1 // 0 + 1
			Green            // 1 + 1
			Blue             // 2 + 1
		)
	*/

	const (
		Red   = (iota * 2) + 1 // (0 * 2) + 1
		Green                  // (1 * 2) + 1
		Blue                   // (2 * 2)+ 1
	)
	fmt.Printf("Red=%d, Green=%d and Blue=%d\n", Red, Green, Blue)

	// Mimicking Unix file permissions
	const (
		X = iota << 1
		W
		R
		XW = X | W
	)
	// fmt.Println(X, W, R)
	fmt.Printf("%b %b %b\n", X, W, R)
	fmt.Printf("%b\n", XW)
}
