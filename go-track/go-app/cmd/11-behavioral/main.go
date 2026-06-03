// Command 11-behavioral exercises Strategy, Observer, Command, Template Method,
// and State.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/behavioral"
	"oop-patterns-learning/go-app/internal/money"
)

func main() {
	// Strategy as a function value: pick a shipping algorithm at runtime.
	strategies := map[string]behavioral.ShippingCost{
		"flat":   behavioral.FlatRate(money.FromMajor(5, 0, "USD")),
		"weight": behavioral.ByWeight(money.FromMajor(2, 0, "USD")),
	}
	fmt.Println("strategy:")
	for _, name := range []string{"flat", "weight"} { // fixed order; map iteration is random
		fmt.Printf("  %-7s 1500g -> %s\n", name, strategies[name](1500))
	}

	// Observer: two subscribers react to one published event.
	pub := &behavioral.Publisher{}
	pub.Subscribe(func(e behavioral.Event) { fmt.Printf("  [email] order %s: %s\n", e.OrderID, e.Name) })
	pub.Subscribe(func(e behavioral.Event) { fmt.Printf("  [audit] order %s: %s\n", e.OrderID, e.Name) })
	fmt.Println("observer:")
	pub.Publish(behavioral.Event{Name: "placed", OrderID: "A-100"})

	// Command: run with full undo support.
	ledger := &behavioral.Ledger{}
	inv := &behavioral.Invoker{}
	inv.Run(behavioral.Deposit{L: ledger, Amount: 100})
	inv.Run(behavioral.Deposit{L: ledger, Amount: 50})
	fmt.Printf("command: balance after 2 deposits = %d\n", ledger.Balance)
	inv.UndoLast()
	fmt.Printf("command: balance after undo       = %d\n", ledger.Balance)

	// Template Method: fixed frame, varying body.
	fmt.Println("template:")
	fmt.Println(behavioral.RenderReport("Daily Sales", func() string { return "  units: 42\n  revenue: $1,337" }))

	// State: legal transitions only.
	fmt.Println("state:")
	var s behavioral.OrderState = behavioral.Pending{}
	s = step(s, "ship") // illegal: unpaid
	s = step(s, "pay")  // pending -> paid
	s = step(s, "ship") // paid -> shipped
	s = step(s, "pay")  // illegal: already shipped
	_ = s
}

func step(s behavioral.OrderState, action string) behavioral.OrderState {
	var next behavioral.OrderState
	var err error
	switch action {
	case "pay":
		next, err = s.Pay()
	case "ship":
		next, err = s.Ship()
	}
	if err != nil {
		fmt.Printf("  %-7s from %-8s -> rejected (%v)\n", action, s.Name(), err)
		return s
	}
	fmt.Printf("  %-7s from %-8s -> %s\n", action, s.Name(), next.Name())
	return next
}
