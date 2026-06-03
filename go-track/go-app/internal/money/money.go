// Package money is a small, immutable value type for amounts of money.
//
// It is the running example for encapsulation (lesson 02) and reappears as the
// canonical "value object" in lesson 16. Two design choices to notice:
//
//   - Amounts are stored as int64 minor units (cents), never float64 — money and
//     floating point do not mix (0.1 + 0.2 != 0.3).
//   - The fields are unexported, so the only way to build a Money is through a
//     constructor, and the only way to "change" one is to derive a new one.
//     That is encapsulation protecting an invariant.
package money

import (
	"errors"
	"fmt"
)

// ErrCurrencyMismatch is returned when an operation mixes two currencies.
var ErrCurrencyMismatch = errors.New("money: currency mismatch")

// Money is an amount in a single currency. It is a value type: copy it freely,
// compare it with Equals, and treat every instance as immutable.
type Money struct {
	minor    int64  // amount in minor units, e.g. cents
	currency string // ISO 4217 code, e.g. "USD"
}

// New builds Money from a raw minor-unit amount (cents).
func New(minorUnits int64, currency string) Money {
	return Money{minor: minorUnits, currency: currency}
}

// FromMajor builds Money from major + minor units, e.g. FromMajor(19, 99, "USD") == $19.99.
func FromMajor(units, cents int64, currency string) Money {
	return Money{minor: units*100 + cents, currency: currency}
}

// Amount returns the value in minor units (cents).
func (m Money) Amount() int64 { return m.minor }

// Currency returns the ISO 4217 currency code.
func (m Money) Currency() string { return m.currency }

// IsZero reports whether the amount is zero.
func (m Money) IsZero() bool { return m.minor == 0 }

// Add returns a NEW Money; the receiver is never mutated.
func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, fmt.Errorf("%w: %s + %s", ErrCurrencyMismatch, m.currency, other.currency)
	}
	return Money{minor: m.minor + other.minor, currency: m.currency}, nil
}

// Sub returns a new Money equal to m minus other.
func (m Money) Sub(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, fmt.Errorf("%w: %s - %s", ErrCurrencyMismatch, m.currency, other.currency)
	}
	return Money{minor: m.minor - other.minor, currency: m.currency}, nil
}

// MultiplyInt scales the amount, e.g. unit price × quantity.
func (m Money) MultiplyInt(factor int64) Money {
	return Money{minor: m.minor * factor, currency: m.currency}
}

// Percentage returns the given whole-number percentage of the amount,
// rounded to the nearest minor unit. Assumes a non-negative amount.
func (m Money) Percentage(percent int64) Money {
	return Money{minor: (m.minor*percent + 50) / 100, currency: m.currency}
}

// Equals compares by value (amount + currency), not by identity.
func (m Money) Equals(other Money) bool {
	return m.minor == other.minor && m.currency == other.currency
}

// String renders the amount for humans, e.g. "$19.99".
func (m Money) String() string {
	a, sign := m.minor, ""
	if a < 0 {
		a, sign = -a, "-"
	}
	return fmt.Sprintf("%s%s%d.%02d", sign, symbol(m.currency), a/100, a%100)
}

func symbol(currency string) string {
	switch currency {
	case "USD":
		return "$"
	case "EUR":
		return "€"
	case "GBP":
		return "£"
	default:
		return currency + " "
	}
}
