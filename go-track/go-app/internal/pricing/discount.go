// Package pricing models discounts behind a single interface. This is the seed
// of the Strategy pattern (lesson 11): interchangeable behaviors hidden behind
// one type, selected at runtime.
package pricing

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/money"
)

// Discount is the abstraction every discount rule satisfies.
type Discount interface {
	// Apply returns the price after the discount.
	Apply(price money.Money) money.Money
	// Label is a human-readable description.
	Label() string
}

// NoDiscount leaves the price untouched (a useful "null object", see lesson 10).
type NoDiscount struct{}

func (NoDiscount) Apply(price money.Money) money.Money { return price }
func (NoDiscount) Label() string                       { return "no discount" }

// PercentageOff takes a whole-number percentage off the price.
type PercentageOff struct{ Percent int64 }

func (d PercentageOff) Apply(price money.Money) money.Money {
	off := price.Percentage(d.Percent)
	return money.New(price.Amount()-off.Amount(), price.Currency())
}

func (d PercentageOff) Label() string { return fmt.Sprintf("%d%% off", d.Percent) }

// FixedAmountOff subtracts a fixed amount, never going below zero.
type FixedAmountOff struct{ Amount money.Money }

func (d FixedAmountOff) Apply(price money.Money) money.Money {
	result := price.Amount() - d.Amount.Amount()
	if result < 0 {
		result = 0
	}
	return money.New(result, price.Currency())
}

func (d FixedAmountOff) Label() string { return d.Amount.String() + " off" }
