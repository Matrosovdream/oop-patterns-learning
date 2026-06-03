// Command 17-aggregates-events shows the Order aggregate ENFORCING its own
// invariants (you can only change it through its methods) and recording domain
// events as state changes.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/ordering/domain"
)

func main() {
	o := domain.NewOrder("ORD-1", "USD")
	must(o.AddLine(domain.Line{SKU: "MUG-1", Name: "Enamel Mug", Unit: money.FromMajor(12, 0, "USD"), Quantity: 2}))

	fmt.Println("the aggregate refuses illegal moves (invariants):")
	report("ship before place", o.Ship()) // not paid, not even placed

	must(o.Place())
	report("add line after place", o.AddLine(domain.Line{SKU: "LATE-1"})) // placed orders are frozen
	report("ship before pay", o.Ship())                                   // must pay first

	fmt.Println("\nthe legal path:")
	must(o.Pay())
	must(o.Ship())
	fmt.Printf("  final status = %s, total = %s\n", o.Status(), o.Total())

	fmt.Println("\ndomain events the aggregate recorded:")
	for _, e := range o.PullEvents() {
		fmt.Printf("  - %s\n", e.EventName())
	}
}

func report(label string, err error) {
	if err != nil {
		fmt.Printf("  %-22s -> rejected: %v\n", label, err)
		return
	}
	fmt.Printf("  %-22s -> allowed\n", label)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
