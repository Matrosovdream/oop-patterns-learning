// Command 19-cqrs separates the WRITE side (commands changing aggregates) from
// the READ side (a denormalized model built by projecting events), and ends with
// a taste of event sourcing (rebuilding state from an event stream).
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/ordering/app"
	"oop-patterns-learning/go-app/internal/ordering/domain"
	"oop-patterns-learning/go-app/internal/ordering/infra"
)

func main() {
	orders := infra.NewInMemoryOrders()
	readModel := app.NewSummaryProjection()

	// One stream of events drives BOTH a log and the read-model projection.
	pub := infra.FanOutPublisher{Targets: []app.EventPublisher{infra.LoggingPublisher{}, readModel}}

	place := app.NewPlaceOrderHandler(orders, pub)
	pay := app.NewPayOrderHandler(orders, pub)
	ship := app.NewShipOrderHandler(orders, pub)

	// ---- WRITE side: state changes only ever go through aggregates. ----
	fmt.Println("commands (write side):")
	id1, _ := place.Handle(app.PlaceOrderCommand{OrderID: "ORD-1", Currency: "USD",
		Lines: []app.PlaceOrderLine{{SKU: "MUG-1", Name: "Mug", UnitMinor: 1200, Quantity: 2}}})
	id2, _ := place.Handle(app.PlaceOrderCommand{OrderID: "ORD-2", Currency: "USD",
		Lines: []app.PlaceOrderLine{{SKU: "TEE-1", Name: "Tee", UnitMinor: 2500, Quantity: 1}}})
	must(pay.Handle(string(id1)))
	must(ship.Handle(string(id1)))
	_ = id2 // left pending on purpose

	// ---- READ side: query the denormalized model; it never touches an aggregate. ----
	fmt.Println("\nqueries (read side) — projected from the very same events:")
	for _, s := range readModel.All() {
		fmt.Printf("  %s  status=%-7s total=%s\n", s.ID, s.Status, s.Total)
	}

	// ---- Event sourcing taste: the events ARE the source of truth. ----
	stream := []domain.Event{
		domain.OrderPlaced{OrderID: "ES-1", Total: money.FromMajor(30, 0, "USD")},
		domain.OrderPaid{OrderID: "ES-1", Amount: money.FromMajor(30, 0, "USD")},
	}
	fmt.Printf("\nevent sourcing: replaying %d events for ES-1 yields status=%q\n", len(stream), replay(stream))
}

// replay rebuilds the current status purely from the event history — no stored
// state, just a fold over events.
func replay(events []domain.Event) domain.Status {
	st := domain.StatusDraft
	for _, e := range events {
		switch e.(type) {
		case domain.OrderPlaced:
			st = domain.StatusPending
		case domain.OrderPaid:
			st = domain.StatusPaid
		case domain.OrderShipped:
			st = domain.StatusShipped
		}
	}
	return st
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
