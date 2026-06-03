// Package creational shows the GoF creational patterns in Go — and, just as
// importantly, how lightweight they become here. Most are a function or a struct
// of functions; only Builder and the occasional Singleton earn their keep.
package creational

import (
	"fmt"
	"sync"

	"oop-patterns-learning/go-app/internal/cart"
	"oop-patterns-learning/go-app/internal/money"
)

// ---------------------------------------------------------------------------
// Factory Method — "let a function decide the concrete type for me."
// In Go this is just a constructor function returning an interface.
// ---------------------------------------------------------------------------

// Payment is the product type the factory hands back.
type Payment interface {
	Charge(amount money.Money) string
}

type cardPayment struct{}

func (cardPayment) Charge(a money.Money) string { return "charged " + a.String() + " to card" }

type cashPayment struct{}

func (cashPayment) Charge(a money.Money) string { return "collected " + a.String() + " in cash" }

// NewPayment is the factory. Callers name a kind and receive a Payment; they
// never see cardPayment/cashPayment (both unexported). Add a kind here and no
// caller changes.
func NewPayment(kind string) (Payment, error) {
	switch kind {
	case "card":
		return cardPayment{}, nil
	case "cash":
		return cashPayment{}, nil
	default:
		return nil, fmt.Errorf("creational: unknown payment kind %q", kind)
	}
}

// ---------------------------------------------------------------------------
// Abstract Factory — "create a FAMILY of related objects that must match."
// ---------------------------------------------------------------------------

type Charger interface{ Charge(money.Money) string }
type Refunder interface{ Refund(money.Money) string }

// Gateway is the abstract factory: each concrete gateway produces a charger and a
// refunder that belong together (you can't accidentally pair Stripe's charger
// with PayPal's refunder).
type Gateway interface {
	Charger() Charger
	Refunder() Refunder
	Name() string
}

type stripe struct{}

func (stripe) Name() string       { return "stripe" }
func (stripe) Charger() Charger   { return stripeCharger{} }
func (stripe) Refunder() Refunder { return stripeRefunder{} }

type stripeCharger struct{}

func (stripeCharger) Charge(a money.Money) string { return "stripe charge " + a.String() }

type stripeRefunder struct{}

func (stripeRefunder) Refund(a money.Money) string { return "stripe refund " + a.String() }

// NewGateway selects a family.
func NewGateway(name string) (Gateway, error) {
	switch name {
	case "stripe":
		return stripe{}, nil
	default:
		return nil, fmt.Errorf("creational: unknown gateway %q", name)
	}
}

// ---------------------------------------------------------------------------
// Builder — "assemble a complex object step by step." This one genuinely helps
// in Go when an object has many optional parts. Fluent chaining returns *Builder.
// ---------------------------------------------------------------------------

type CartBuilder struct {
	currency string
	lines    []cart.Line
}

func NewCartBuilder(currency string) *CartBuilder {
	return &CartBuilder{currency: currency}
}

func (b *CartBuilder) Add(sku, name string, unit money.Money, qty int64) *CartBuilder {
	b.lines = append(b.lines, cart.Line{SKU: sku, Name: name, Unit: unit, Quantity: qty})
	return b
}

func (b *CartBuilder) Build() *cart.Cart {
	c := cart.New(b.currency)
	for _, l := range b.lines {
		c.Add(l)
	}
	return c
}

// ---------------------------------------------------------------------------
// Singleton — "exactly one instance, created once." In Go: a package var, lazily
// initialized with sync.Once (concurrency-safe). Use sparingly; a singleton is a
// global, and globals fight dependency injection.
// ---------------------------------------------------------------------------

type Config struct {
	Currency   string
	TaxPercent int64
}

var (
	cfgOnce sync.Once
	cfg     *Config
)

// DefaultConfig returns the one shared Config, building it on first call only.
func DefaultConfig() *Config {
	cfgOnce.Do(func() {
		cfg = &Config{Currency: "USD", TaxPercent: 8}
	})
	return cfg
}
