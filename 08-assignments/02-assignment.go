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
	primes := generatePrimes(start, end)
	for _, primeNo := range primes {
		fmt.Printf("Prime : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) []int {
	var primes []int
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
