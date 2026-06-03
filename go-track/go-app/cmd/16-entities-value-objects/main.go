// Command 16-entities-value-objects shows the single most useful distinction in
// tactical DDD: equality by VALUE (value objects) vs equality by IDENTITY
// (entities).
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/ordering/domain"
)

func main() {
	// VALUE OBJECTS — equal by their contents; interchangeable. Two $10.00 USD
	// are "the same" the way two $10 bills are the same.
	a := money.FromMajor(10, 0, "USD")
	b := money.FromMajor(10, 0, "USD")
	fmt.Printf("value object equality:  %s == %s ? %v\n", a, b, a.Equals(b))

	// A Line is a value object too: equal by SKU + price + quantity (Go's == works
	// because all its fields are comparable).
	l1 := domain.Line{SKU: "MUG-1", Name: "Mug", Unit: a, Quantity: 1}
	l2 := domain.Line{SKU: "MUG-1", Name: "Mug", Unit: b, Quantity: 1}
	fmt.Printf("line (VO) equality:     %v\n", l1 == l2)

	// ENTITIES — equal by IDENTITY, not contents. Two orders with identical lines
	// are still different orders if their IDs differ; and an order stays the same
	// order even as its contents change over time.
	o1 := domain.NewOrder("ORD-1", "USD")
	o2 := domain.NewOrder("ORD-2", "USD")
	_ = o1.AddLine(l1)
	_ = o2.AddLine(l1) // identical contents…
	fmt.Printf("entity identity:        same id? %v  (identical contents, different identities)\n", o1.ID() == o2.ID())

	fmt.Println("\nRule of thumb: if you'd track it through changes, it's an entity (give it an")
	fmt.Println("ID). If you only care what it IS right now, it's a value object (make it immutable).")
}
