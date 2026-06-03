// Command 18-application-layer shows a use case orchestrating the domain, and an
// Anti-Corruption Layer (ACL) translating another context's model into this
// one's DTOs so foreign shapes never leak into the domain.
package main

import (
	"fmt"
	"time"

	"oop-patterns-learning/go-app/internal/catalog"
	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/ordering/app"
	"oop-patterns-learning/go-app/internal/ordering/infra"
)

func main() {
	// The Catalog context hands us products in ITS model (catalog.Product).
	now := time.Now()
	products := []*catalog.Product{
		catalog.NewProduct("MUG-1", "Enamel Mug", money.FromMajor(12, 0, "USD"), now),
		catalog.NewProduct("TEE-1", "Logo Tee", money.FromMajor(25, 0, "USD"), now),
	}

	// ---- Anti-Corruption Layer ----
	// Translate catalog.Product → ordering DTO. The Ordering domain never imports
	// catalog; if Catalog renames a field tomorrow, only this function changes.
	lines := translate(products)

	// ---- Wire + run the use case ----
	orders := infra.NewInMemoryOrders()
	place := app.NewPlaceOrderHandler(orders, infra.LoggingPublisher{})

	id, err := place.Handle(app.PlaceOrderCommand{OrderID: "ORD-1", Currency: "USD", Lines: lines})
	if err != nil {
		panic(err)
	}

	o, _ := orders.Find(id)
	fmt.Printf("\nplaced order %s from %d catalog products, total=%s\n", o.ID(), len(products), o.Total())
	fmt.Println("The use case held no business rules — it only choreographed domain + ports.")
}

func translate(products []*catalog.Product) []app.PlaceOrderLine {
	out := make([]app.PlaceOrderLine, 0, len(products))
	for _, p := range products {
		out = append(out, app.PlaceOrderLine{
			SKU:       p.SKU,
			Name:      p.Name,
			UnitMinor: p.Price.Amount(), // pull the minor units out of the value object
			Quantity:  1,
		})
	}
	return out
}
