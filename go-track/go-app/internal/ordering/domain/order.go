// Package domain is the heart of the Ordering bounded context. It is the
// innermost layer of the architecture: it imports only money (another pure
// domain type) and the standard library — never a database, web framework, or
// the application layer. Everything else depends on it; it depends on nothing
// outward. That is the dependency rule (lesson 14) made concrete.
package domain

import (
	"errors"

	"oop-patterns-learning/go-app/internal/money"
)

// Status is the order's place in its lifecycle.
type Status string

const (
	StatusDraft   Status = "draft"   // being assembled; lines can be added
	StatusPending Status = "pending" // placed, awaiting payment
	StatusPaid    Status = "paid"    // paid, awaiting shipment
	StatusShipped Status = "shipped" // done
)

// Domain errors — invariant violations expressed in the domain's own language.
var (
	ErrEmptyOrder       = errors.New("ordering: cannot place an order with no lines")
	ErrModifyAfterPlace = errors.New("ordering: cannot change an order once it is placed")
	ErrNotPayable       = errors.New("ordering: order cannot be paid from its current state")
	ErrNotShippable     = errors.New("ordering: order cannot be shipped from its current state")
)

// OrderID is the aggregate's identity (a value object). Two orders are the same
// order iff their IDs match — not because their contents match.
type OrderID string

// Line is a line item: a value object living INSIDE the Order aggregate. It has
// no identity of its own and is only ever reached through its Order.
type Line struct {
	SKU      string
	Name     string
	Unit     money.Money
	Quantity int64
}

// Subtotal is unit price × quantity.
func (l Line) Subtotal() money.Money { return l.Unit.MultiplyInt(l.Quantity) }

// Order is the AGGREGATE ROOT. Outside code may only change it through these
// methods, which is what keeps its invariants ("a placed order is immutable",
// "you can't ship before paying") always true. State is unexported precisely so
// nothing can bypass the rules.
type Order struct {
	id       OrderID
	currency string
	lines    []Line
	status   Status
	events   []Event
}

// NewOrder starts a draft order.
func NewOrder(id OrderID, currency string) *Order {
	return &Order{id: id, currency: currency, status: StatusDraft}
}

// AddLine adds an item — only while the order is still a draft.
func (o *Order) AddLine(l Line) error {
	if o.status != StatusDraft {
		return ErrModifyAfterPlace
	}
	o.lines = append(o.lines, l)
	return nil
}

// Place confirms the order: it must be a non-empty draft. Emits OrderPlaced.
func (o *Order) Place() error {
	if o.status != StatusDraft {
		return ErrModifyAfterPlace
	}
	if len(o.lines) == 0 {
		return ErrEmptyOrder
	}
	o.status = StatusPending
	o.record(OrderPlaced{OrderID: string(o.id), Total: o.Total()})
	return nil
}

// Pay moves pending → paid. Emits OrderPaid.
func (o *Order) Pay() error {
	if o.status != StatusPending {
		return ErrNotPayable
	}
	o.status = StatusPaid
	o.record(OrderPaid{OrderID: string(o.id), Amount: o.Total()})
	return nil
}

// Ship moves paid → shipped. Emits OrderShipped.
func (o *Order) Ship() error {
	if o.status != StatusPaid {
		return ErrNotShippable
	}
	o.status = StatusShipped
	o.record(OrderShipped{OrderID: string(o.id)})
	return nil
}

// Total sums the lines. The aggregate owns this computation so the total is
// always consistent with the lines.
func (o *Order) Total() money.Money {
	total := money.New(0, o.currency)
	for _, l := range o.lines {
		total, _ = total.Add(l.Subtotal())
	}
	return total
}

// ID returns the identity.
func (o *Order) ID() OrderID { return o.id }

// Status returns the current lifecycle state.
func (o *Order) Status() Status { return o.status }

// Lines returns a copy so callers can't mutate the aggregate's internals.
func (o *Order) Lines() []Line { return append([]Line(nil), o.lines...) }

// PullEvents returns the events recorded since the last pull and clears them.
// The application layer pulls these after saving and hands them to a publisher.
func (o *Order) PullEvents() []Event {
	evs := o.events
	o.events = nil
	return evs
}

func (o *Order) record(e Event) { o.events = append(o.events, e) }
