package main

import "fmt"

func main() {
	var productRanks map[string]int
	/*
		productRanks = make(map[string]int)
		productRanks["pen"] = 4
		productRanks["pencil"] = 1
		productRanks["marker"] = 5
	*/
	// productRanks = map[string]int{"pen": 4, "pencil": 1, "marker": 5}
	productRanks = map[string]int{
		"pen":    4,
		"pencil": 1,
		"marker": 5,
	}
	fmt.Println(productRanks)

	fmt.Println()
	fmt.Printf("len(productRanks) = %d\n", len(productRanks))

	fmt.Println()
	fmt.Println("Iterating using range")
	for product, rank := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", product, rank)
	}

	fmt.Println()
	keyToCheck := "notebook"
	fmt.Printf("Check for the existance of key [%q]\n", keyToCheck)
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("productRanks[%q] = %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key [%q] does not exist!\n", keyToCheck)
	}

	fmt.Println()
	keyToRemove := "notepad"
	fmt.Printf("Removing a key [%q]\n", keyToRemove)
	delete(productRanks, keyToRemove)
	fmt.Println(productRanks)
}
