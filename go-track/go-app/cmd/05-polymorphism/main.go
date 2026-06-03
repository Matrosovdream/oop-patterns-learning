// Command 05-polymorphism shows runtime polymorphism the Go way: a slice of an
// interface type, each element a different concrete implementation, all treated
// uniformly. No inheritance, no virtual-method tables you manage by hand.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/pricing"
)

func main() {
	listPrice := money.FromMajor(80, 00, "USD")

	// Each of these satisfies pricing.Discount implicitly. We can hold them all
	// in one []pricing.Discount and call Apply without caring which is which.
	discounts := []pricing.Discount{
		pricing.NoDiscount{},
		pricing.PercentageOff{Percent: 10},
		pricing.PercentageOff{Percent: 25},
		pricing.FixedAmountOff{Amount: money.FromMajor(15, 00, "USD")},
	}

	fmt.Printf("list price: %s\n\n", listPrice)
	for _, d := range discounts {
		fmt.Printf("%-12s -> %s\n", d.Label(), d.Apply(listPrice))
	}

	fmt.Println("\nOne loop, many behaviors — polymorphism through interfaces.")
}
