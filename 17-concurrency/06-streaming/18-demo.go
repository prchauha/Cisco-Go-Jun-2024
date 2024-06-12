package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("[main] All data received!")
}

func genNos(ch chan int) {
	count := rand.Intn(20)
	fmt.Println("[genNos] count :", count)
	for i := range count {
		ch <- 10 * (i + 1)
	}
	fmt.Println("[genNos] closing the channel")
	close(ch)
}
