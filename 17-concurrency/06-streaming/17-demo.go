package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for {
		if data, open := <-ch; open {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
			continue
		}
		fmt.Println("[main] All data received!")
		break
	}
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
