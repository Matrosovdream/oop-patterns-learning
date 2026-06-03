// Package cart is the ShopKit shopping cart: a list of line items that can total
// itself. It's intentionally small — its single responsibility is "hold lines and
// compute a subtotal", nothing about discounts, tax, persistence, or display.
package cart

import "oop-patterns-learning/go-app/internal/money"

// Line is one row in the cart.
type Line struct {
	SKU      string
	Name     string
	Unit     money.Money
	Quantity int64
}

// Subtotal is unit price × quantity.
func (l Line) Subtotal() money.Money { return l.Unit.MultiplyInt(l.Quantity) }

// Cart collects lines in a single currency.
type Cart struct {
	currency string
	lines    []Line
}

// New creates an empty cart for the given currency.
func New(currency string) *Cart { return &Cart{currency: currency} }

// Add appends a line.
func (c *Cart) Add(l Line) { c.lines = append(c.lines, l) }

// Lines returns the cart's lines (read-only use intended).
func (c *Cart) Lines() []Line { return c.lines }

// Currency reports the cart's currency.
func (c *Cart) Currency() string { return c.currency }

// Subtotal sums every line. Same-currency adds never error here.
func (c *Cart) Subtotal() money.Money {
	total := money.New(0, c.currency)
	for _, l := range c.lines {
		total, _ = total.Add(l.Subtotal())
	}
	return total
}
