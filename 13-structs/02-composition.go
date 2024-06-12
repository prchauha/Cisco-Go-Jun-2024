package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	// Dummy
	Product
	Expiry string
}

func main() {
	// var grapes PerishableProduct
	grapes := PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Grapes",
			Cost: 50,
		},
		Expiry: "2 Days",
	}
	fmt.Printf("%#v\n", grapes)
	// fmt.Println(grapes.Product.Cost)
	// fmt.Println(grapes.Cost)

	// Use the Format & ApplyDiscount functions with 'grapes'
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercent float64) {
	p.Cost = p.Cost * ((100 - discountPercent) / 100)
}
