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

// Factory function to hide the complexity of creating a perishable product
func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
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
	fmt.Println(Format(grapes.Product))
	fmt.Println("After applying 10% discount")
	ApplyDiscount(&grapes.Product, 10)
	fmt.Println(Format(grapes.Product))

	var milk = NewPerishableProduct(200, "Milk", 40, "1 Day")
	fmt.Printf("%#v\n", milk)
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercent float64) {
	p.Cost = p.Cost * ((100 - discountPercent) / 100)
}
