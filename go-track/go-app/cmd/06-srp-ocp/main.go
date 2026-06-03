// Command 06-srp-ocp demonstrates two principles at once:
//
//   - SRP: each piece has ONE reason to change. The cart totals; the discount
//     decides a reduction; formatting renders text; sending delivers it. Mixing
//     these into one "OrderService" would give it four reasons to change.
//   - OCP: Checkout is OPEN to new discounts but CLOSED to modification — adding a
//     discount type never edits this function, because it depends on the
//     pricing.Discount abstraction, not a switch over discount kinds.
package main

import (
	"fmt"
	"strings"

	"oop-patterns-learning/go-app/internal/cart"
	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/notify"
	"oop-patterns-learning/go-app/internal/pricing"
)

// Checkout computes what the customer pays. Notice it takes a Discount — any
// discount, including ones written after this function. That's Open/Closed.
func Checkout(c *cart.Cart, d pricing.Discount) (subtotal, total money.Money) {
	subtotal = c.Subtotal()
	total = d.Apply(subtotal)
	return subtotal, total
}

// renderReceipt has the single responsibility of formatting. It returns a string;
// it does not print, email, or save. That keeps it trivial to reuse and change.
func renderReceipt(c *cart.Cart, d pricing.Discount, subtotal, total money.Money) string {
	var b strings.Builder
	b.WriteString("ShopKit receipt\n")
	for _, l := range c.Lines() {
		fmt.Fprintf(&b, "  %-14s %d × %s = %s\n", l.Name, l.Quantity, l.Unit, l.Subtotal())
	}
	fmt.Fprintf(&b, "  subtotal:     %s\n", subtotal)
	fmt.Fprintf(&b, "  discount:     %s\n", d.Label())
	fmt.Fprintf(&b, "  total:        %s\n", total)
	return b.String()
}

func main() {
	c := cart.New("USD")
	c.Add(cart.Line{SKU: "MUG-1", Name: "Enamel Mug", Unit: money.FromMajor(12, 0, "USD"), Quantity: 2})
	c.Add(cart.Line{SKU: "TEE-1", Name: "Logo Tee", Unit: money.FromMajor(25, 0, "USD"), Quantity: 1})

	// OCP: swap the discount freely; Checkout is untouched.
	discount := pricing.PercentageOff{Percent: 10}
	subtotal, total := Checkout(c, discount)

	// SRP: format, then (separately) deliver. Different jobs, different code.
	receipt := renderReceipt(c, discount, subtotal, total)
	sender := notify.EmailSender{From: "orders@shopkit.test"}
	_ = sender.Send("buyer@example.com", receipt)
}
