/*
Refactor the logic for generating prime nos to a "generatePrimes" function
Print the prime numbers in the main() function
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
NO_LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue NO_LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}
}
