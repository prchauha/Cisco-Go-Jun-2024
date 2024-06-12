package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Aliquip culpa ullamco reprehenderit tempor ut quis nisi consectetur aute reprehenderit elit."
	x = 99.99
	x = true
	x = Product{100, "Pen", 10}
	fmt.Println(x)

	// treating x as an int
	x = 200
	// x = "Ullamco consectetur dolore exercitation cupidatat consequat consequat est ullamco et fugiat ut sunt labore."
	// unsafe conversion
	y := x.(int) + 300
	fmt.Println(y)

	// safe conversion through "type assertion"
	x = "Deserunt qui eu exercitation veniam ex ex aliqua duis."
	if val, ok := x.(int); ok {
		y := val + 300
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// using "type switch" for "type assertion"
	// x = 100
	// x = "Aliquip culpa ullamco reprehenderit tempor ut quis nisi consectetur aute reprehenderit elit."
	// x = 99.99
	// x = true
	x = Product{100, "Pen", 10}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 300 =", val+300)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case float64:
		fmt.Println("x is a float64, 0.01 * x =", val*0.01)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	default:
		fmt.Println("x is an unknown type")
	}

}
