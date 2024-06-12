/* Using a custom type to encapsulate the concurrent safe data manipulation*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var c Counter

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", c.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.Increment()
}
