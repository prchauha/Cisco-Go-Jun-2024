package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tkmagesh/cisco-go-jun-2024/12-modularity-app/calculator"
	ut "github.com/tkmagesh/cisco-go-jun-2024/12-modularity-app/calculator/utils"
)

func main() {
	color.Red("Welcome!")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.OpCount())
	fmt.Println(ut.IsEven(21))
	run()
}
