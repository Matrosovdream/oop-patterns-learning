// Command 12-layered shows the classic layered flow of one request:
//
//	presentation (this main) → application (use case) → domain → infrastructure
//
// The golden rule: dependencies point DOWN toward the domain. The domain never
// imports the application or infrastructure.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/ordering/app"
	"oop-patterns-learning/go-app/internal/ordering/infra"
)

func main() {
	// ---- Composition root: build the layers bottom-up and wire them. ----
	orders := infra.NewInMemoryOrders()                       // infrastructure (adapter)
	publisher := infra.LoggingPublisher{}                     // infrastructure (adapter)
	placeOrder := app.NewPlaceOrderHandler(orders, publisher) // application (use case)

	// ---- Presentation: turn an inbound "request" into a command, call the use case. ----
	cmd := app.PlaceOrderCommand{
		OrderID:  "ORD-1",
		Currency: "USD",
		Lines: []app.PlaceOrderLine{
			{SKU: "MUG-1", Name: "Enamel Mug", UnitMinor: 1200, Quantity: 2},
			{SKU: "TEE-1", Name: "Logo Tee", UnitMinor: 2500, Quantity: 1},
		},
	}

	fmt.Println("placing order (events emitted by the domain):")
	id, err := placeOrder.Handle(cmd)
	if err != nil {
		panic(err)
	}

	// ---- Read back a domain object through the repository. ----
	o, err := orders.Find(id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\norder %s — status=%s, total=%s\n", o.ID(), o.Status(), o.Total())
	fmt.Println("\nNote: domain/order.go imports neither app nor infra — only money.")
}
