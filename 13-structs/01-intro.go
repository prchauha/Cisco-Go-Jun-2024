package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/*
		var pen Product
		pen.Id = 100
		pen.Name = "Pen"
		pen.Cost = 10
	*/
	// fmt.Println(pen)
	// fmt.Printf("%#v\n", pen)
	// var pen Product = Product{100, "Pen", 10}
	// var pen Product = Product{Id: 100, Name: "Pen", Cost: 10}
	/*
		var pen Product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/

	// type inference
	/*
		var pen = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/

	// using :=
	pen := Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	fmt.Printf("%+v\n", pen)

	// creates a copy
	pen2 := pen
	pen2.Cost = 100
	fmt.Printf("pen.Cost = %v, pen2.Cost = %v\n", pen.Cost, pen2.Cost)

	fmt.Println(Format(pen))
	fmt.Println("After applying 10% discount...")
	ApplyDiscount(&pen, 10)
	fmt.Println(Format(pen))
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercent float64) {
	p.Cost = p.Cost * ((100 - discountPercent) / 100)
}
