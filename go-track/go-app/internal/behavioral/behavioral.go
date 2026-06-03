// Package behavioral shows the GoF behavioral patterns. Several of these are so
// light in Go that they're "just a function value" — that's the lesson.
package behavioral

import (
	"errors"
	"fmt"

	"oop-patterns-learning/go-app/internal/money"
)

// ---------------------------------------------------------------------------
// Strategy — a swappable algorithm. The Go-idiomatic form is a function type,
// not an interface with one method. (pricing.Discount from lesson 05 is the
// interface form; this is the function form.)
// ---------------------------------------------------------------------------

// ShippingCost computes shipping for a parcel weight. Swap the algorithm by
// swapping the function — no class hierarchy required.
type ShippingCost func(weightGrams int64) money.Money

func FlatRate(amount money.Money) ShippingCost {
	return func(int64) money.Money { return amount }
}

func ByWeight(perKg money.Money) ShippingCost {
	return func(g int64) money.Money {
		kg := (g + 999) / 1000 // round up to the next kg
		return perKg.MultiplyInt(kg)
	}
}

// ---------------------------------------------------------------------------
// Observer — publishers notify subscribers without knowing who they are.
// ---------------------------------------------------------------------------

type Event struct{ Name, OrderID string }

// Observer reacts to an event. Using a func keeps subscribers trivial to add.
type Observer func(Event)

type Publisher struct{ subs []Observer }

func (p *Publisher) Subscribe(o Observer) { p.subs = append(p.subs, o) }

func (p *Publisher) Publish(e Event) {
	for _, o := range p.subs {
		o(e)
	}
}

// ---------------------------------------------------------------------------
// Command — package an action (and its undo) as a value, so you can queue,
// log, or reverse it.
// ---------------------------------------------------------------------------

// Ledger is the receiver the commands act on.
type Ledger struct{ Balance int64 }

type Command interface {
	Execute()
	Undo()
}

type Deposit struct {
	L      *Ledger
	Amount int64
}

func (d Deposit) Execute() { d.L.Balance += d.Amount }
func (d Deposit) Undo()    { d.L.Balance -= d.Amount }

// Invoker runs commands and can undo the most recent one.
type Invoker struct{ history []Command }

func (in *Invoker) Run(c Command) {
	c.Execute()
	in.history = append(in.history, c)
}

func (in *Invoker) UndoLast() {
	if len(in.history) == 0 {
		return
	}
	last := in.history[len(in.history)-1]
	in.history = in.history[:len(in.history)-1]
	last.Undo()
}

// ---------------------------------------------------------------------------
// Template Method — fix the skeleton of an algorithm, vary specific steps.
// Classic OOP uses an abstract method overridden by subclasses; in Go the
// varying step is a function parameter.
// ---------------------------------------------------------------------------

func RenderReport(title string, body func() string) string {
	return fmt.Sprintf("== %s ==\n%s\n-- end --", title, body())
}

// ---------------------------------------------------------------------------
// State — behavior changes with internal state; each state is its own type that
// knows which transitions are legal.
// ---------------------------------------------------------------------------

var ErrIllegalTransition = errors.New("behavioral: illegal state transition")

// OrderState is a state in the order lifecycle. Transitions return the next
// state (or an error if the move isn't allowed from here).
type OrderState interface {
	Pay() (OrderState, error)
	Ship() (OrderState, error)
	Name() string
}

type Pending struct{}

func (Pending) Name() string              { return "pending" }
func (Pending) Pay() (OrderState, error)  { return Paid{}, nil }
func (Pending) Ship() (OrderState, error) { return nil, ErrIllegalTransition } // can't ship unpaid

type Paid struct{}

func (Paid) Name() string              { return "paid" }
func (Paid) Pay() (OrderState, error)  { return nil, ErrIllegalTransition } // already paid
func (Paid) Ship() (OrderState, error) { return Shipped{}, nil }

type Shipped struct{}

func (Shipped) Name() string              { return "shipped" }
func (Shipped) Pay() (OrderState, error)  { return nil, ErrIllegalTransition }
func (Shipped) Ship() (OrderState, error) { return nil, ErrIllegalTransition }
