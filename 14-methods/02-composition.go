package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercent float64) {
	p.Cost = p.Cost * ((100 - discountPercent) / 100)
}

type PerishableProduct struct {
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

// method overriding
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
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
	fmt.Println(grapes.Cost)

	// Use the Format & ApplyDiscount "functions" with 'grapes'
	/*
		fmt.Println(Format(grapes.Product))
		fmt.Println("After applying 10% discount")
		ApplyDiscount(&grapes.Product, 10)
		fmt.Println(Format(grapes.Product))
	*/

	// Use the Format & ApplyDiscount "methods" with 'grapes'
	fmt.Println(grapes.Format())
	fmt.Println("After applying 10% discount")
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())

}
