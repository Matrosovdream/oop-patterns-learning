// Command 13-hexagonal shows ports & adapters: the same use case driven through
// the same PORT (app.EventPublisher) by two different ADAPTERS, swapped with one
// line, the application code untouched.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/ordering/app"
	"oop-patterns-learning/go-app/internal/ordering/domain"
	"oop-patterns-learning/go-app/internal/ordering/infra"
)

// recorder is a second adapter for the app.EventPublisher port, written right
// here. It captures event names instead of printing them — exactly what you'd do
// in a test. The application has no idea which adapter it's talking to.
type recorder struct{ names []string }

func (r *recorder) Publish(events []domain.Event) {
	for _, e := range events {
		r.names = append(r.names, e.EventName())
	}
}

func placeSample(publisher app.EventPublisher, id string) {
	orders := infra.NewInMemoryOrders()
	handler := app.NewPlaceOrderHandler(orders, publisher) // <- same handler, any adapter
	_, err := handler.Handle(app.PlaceOrderCommand{
		OrderID:  id,
		Currency: "USD",
		Lines:    []app.PlaceOrderLine{{SKU: "MUG-1", Name: "Mug", UnitMinor: 1200, Quantity: 1}},
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	// Adapter A: the logging adapter (a "real" side effect).
	fmt.Println("driving the port with the LOGGING adapter:")
	placeSample(infra.LoggingPublisher{}, "ORD-A")

	// Adapter B: the recording adapter (a test double) — swapped in trivially.
	rec := &recorder{}
	placeSample(rec, "ORD-B")
	fmt.Printf("\ndriving the port with the RECORDING adapter captured: %v\n", rec.names)

	fmt.Println("\nThe hexagon: app talks to ports (interfaces); adapters plug in from outside.")
}
