package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos []int = []int{3, 1, 4}
	// var nos = []int{3, 1, 4, 2, 5}
	// nos := []int{3, 1, 4, 2, 5}

	var nos []int
	// nos[0] = 3
	nos = append(nos, 3)
	nos = append(nos, 1)
	/*
		nos = append(nos, 4)
		nos = append(nos, 2)
		nos = append(nos, 5)
	*/
	nos = append(nos, 4, 2, 5)

	// appending another slize
	tens := []int{10, 20, 30}
	nos = append(nos, tens...)
	fmt.Printf("nos = %v\n", nos)

	nos2 := nos
	nos2[0] = 9999
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	fmt.Println("Before sorting, nos = ", nos)
	sort(nos)
	fmt.Println("After sorting, nos = ", nos)

	// slicing
	fmt.Println("slicing")
	s2 := nos[2:5]
	fmt.Printf("nos[2:5] = %v\n", s2)

	fmt.Printf("nos[2:] = %v\n", nos[2:])
	fmt.Printf("nos[:5] = %v\n", nos[:5])

}

func sort(list []int) /* no return results */ {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
