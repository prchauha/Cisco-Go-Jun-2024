package main

import (
	"fmt"

	"github.com/tkmagesh/cisco-go-jun-2024/12-modularity-app/calculator"
)

func main() {
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	run()
}
