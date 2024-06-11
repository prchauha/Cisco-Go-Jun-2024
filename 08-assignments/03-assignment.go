/* Find the highest occurring "word size" in the given string and print the "word size" & the number of occurances */
/*
	ex output:
	5 letter words occurs the most with 31 occurances
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Commodo Lorem enim dolore culpa minim ipsum ex excepteur in Commodo duis nulla ex laborum irure sunt incididunt Incididunt amet Lorem amet dolor sit consectetur culpa esse quis laborum pariatur laborum fugiat mollit Mollit voluptate aliquip Lorem incididunt mollit pariatur eu enim proident culpa esse laborum voluptate Nostrud aliqua magna ipsum qui duis euNisi pariatur sit do magna Lorem nostrud voluptate occaecat occaecat quis dolore irure Velit aliqua reprehenderit eu duis aliqua excepteur duis non enim Nostrud qui voluptate enim eiusmod dolor proident laboris nostrud commodo laborum aliquip sunt Exercitation do anim do ullamco Ipsum eiusmod aute qui ea consectetur Veniam laborum occaecat mollit pariatur commodo id ullamco dolore ipsum sit dolore elit fugiatUt ad aliqua dolor quis velit reprehenderit Dolore aliquip exercitation do fugiat Pariatur irure aliqua magna quis"
	words := strings.Split(str, " ")
	wordsCountBySize := getWordsCountBySize(words)
	maxCount, maxSize := getMaxByCount(wordsCountBySize)
	fmt.Printf("%d letter words occurs the most with %d occurrances\n", maxSize, maxCount)
}

func getWordsCountBySize(words []string) map[int]int {
	wordsCountBySize := make(map[int]int)
	for _, word := range words {
		size := len(word)
		wordsCountBySize[size]++
	}
	return wordsCountBySize
}

func getMaxByCount(wordsCountBySize map[int]int) (maxCount, maxSize int) {
	for size, count := range wordsCountBySize {
		if count > maxCount {
			maxCount, maxSize = count, size
		}
	}
	return
}
