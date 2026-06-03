// Command 09-creational exercises Factory Method, Abstract Factory, Builder, and
// Singleton — and shows how small they are in Go.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/creational"
	"oop-patterns-learning/go-app/internal/money"
)

func main() {
	price := money.FromMajor(42, 0, "USD")

	// Factory Method: ask for a kind, get a Payment; concrete type stays hidden.
	for _, kind := range []string{"card", "cash"} {
		p, err := creational.NewPayment(kind)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[factory]  %s\n", p.Charge(price))
	}

	// Abstract Factory: one family yields a matching charger + refunder.
	gw, _ := creational.NewGateway("stripe")
	fmt.Printf("[abstract] %s; %s\n", gw.Charger().Charge(price), gw.Refunder().Refund(price))

	// Builder: assemble a cart fluently, then Build the finished value.
	c := creational.NewCartBuilder("USD").
		Add("MUG-1", "Enamel Mug", money.FromMajor(12, 0, "USD"), 2).
		Add("TEE-1", "Logo Tee", money.FromMajor(25, 0, "USD"), 1).
		Build()
	fmt.Printf("[builder]  cart subtotal %s\n", c.Subtotal())

	// Singleton: same pointer every call.
	a, b := creational.DefaultConfig(), creational.DefaultConfig()
	fmt.Printf("[singleton] same instance? %v (currency=%s, tax=%d%%)\n", a == b, a.Currency, a.TaxPercent)
}
