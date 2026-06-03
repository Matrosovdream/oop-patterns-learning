// Command 14-clean demonstrates the dependency rule of Clean/Onion architecture:
// source dependencies only ever point INWARD, toward the domain. We make that
// rule visible with compile-time interface assertions and a full lifecycle run.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/ordering/app"
	"oop-patterns-learning/go-app/internal/ordering/domain"
	"oop-patterns-learning/go-app/internal/ordering/infra"
)

// Compile-time proof that the OUTER infrastructure adapters satisfy the INNER
// ports. The interfaces (domain.Repository, app.EventPublisher) live in the inner
// layers; the implementations live outside and depend inward. If someone broke
// that direction, these lines would stop compiling.
var (
	_ domain.Repository  = (*infra.InMemoryOrders)(nil)
	_ app.EventPublisher = infra.LoggingPublisher{}
)

func main() {
	orders := infra.NewInMemoryOrders()
	pub := infra.LoggingPublisher{}

	place := app.NewPlaceOrderHandler(orders, pub)
	pay := app.NewPayOrderHandler(orders, pub)
	ship := app.NewShipOrderHandler(orders, pub)

	id, err := place.Handle(app.PlaceOrderCommand{
		OrderID:  "ORD-1",
		Currency: "USD",
		Lines:    []app.PlaceOrderLine{{SKU: "TEE-1", Name: "Logo Tee", UnitMinor: 2500, Quantity: 2}},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("\nrunning the full lifecycle through use cases:")
	must(pay.Handle(string(id)))
	must(ship.Handle(string(id)))

	o, _ := orders.Find(id)
	fmt.Printf("\nfinal: order %s status=%s total=%s\n", o.ID(), o.Status(), o.Total())
	fmt.Println("Dependency rule held — outer layers depend on inner abstractions only.")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
