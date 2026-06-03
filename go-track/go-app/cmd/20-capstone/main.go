// Command 20-capstone wires the whole Ordering bounded context together end to
// end: composition root → use cases → aggregate invariants → domain events →
// CQRS read model. It's everything from lessons 01–19 in one runnable flow.
package main

import (
	"errors"
	"fmt"

	"oop-patterns-learning/go-app/internal/ordering/app"
	"oop-patterns-learning/go-app/internal/ordering/domain"
	"oop-patterns-learning/go-app/internal/ordering/infra"
)

func main() {
	// ---- Composition root: choose adapters, wire ports, build use cases. ----
	orders := infra.NewInMemoryOrders()
	readModel := app.NewSummaryProjection()
	pub := infra.FanOutPublisher{Targets: []app.EventPublisher{infra.LoggingPublisher{}, readModel}}

	place := app.NewPlaceOrderHandler(orders, pub)
	pay := app.NewPayOrderHandler(orders, pub)
	ship := app.NewShipOrderHandler(orders, pub)

	// ---- A happy path: place → pay → ship. ----
	fmt.Println("== ORD-1: full lifecycle ==")
	id, err := place.Handle(app.PlaceOrderCommand{
		OrderID:  "ORD-1",
		Currency: "USD",
		Lines: []app.PlaceOrderLine{
			{SKU: "MUG-1", Name: "Enamel Mug", UnitMinor: 1200, Quantity: 2},
			{SKU: "TEE-1", Name: "Logo Tee", UnitMinor: 2500, Quantity: 1},
		},
	})
	must(err)
	must(pay.Handle(string(id)))
	must(ship.Handle(string(id)))

	// ---- A second order left pending. ----
	fmt.Println("\n== ORD-2: placed, left pending ==")
	_, err = place.Handle(app.PlaceOrderCommand{
		OrderID:  "ORD-2",
		Currency: "USD",
		Lines:    []app.PlaceOrderLine{{SKU: "CAP-1", Name: "Beanie", UnitMinor: 1800, Quantity: 1}},
	})
	must(err)

	// ---- Domain rules still bite: an empty order is rejected. ----
	fmt.Println("\n== ORD-3: invalid (no lines) ==")
	_, err = place.Handle(app.PlaceOrderCommand{OrderID: "ORD-3", Currency: "USD"})
	if errors.Is(err, domain.ErrEmptyOrder) {
		fmt.Println("  correctly rejected: " + err.Error())
	}

	// ---- The read model answers queries without touching aggregates. ----
	fmt.Println("\n== order book (CQRS read model) ==")
	for _, s := range readModel.All() {
		fmt.Printf("  %-6s status=%-7s total=%s\n", s.ID, s.Status, s.Total)
	}

	fmt.Println("\nThat's the whole stack: SOLID + patterns + hexagonal + DDD + CQRS, working together.")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
