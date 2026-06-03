package domain

import "oop-patterns-learning/go-app/internal/money"

// Event is a DOMAIN EVENT: a record that something meaningful happened, named in
// the past tense and in the ubiquitous language. Domain events are how an
// aggregate tells the rest of the system "this occurred" without depending on
// who's listening (see lesson 17, and CQRS projections in lesson 19).
type Event interface {
	EventName() string
}

// OrderPlaced is emitted when a draft order is confirmed.
type OrderPlaced struct {
	OrderID string
	Total   money.Money
}

func (OrderPlaced) EventName() string { return "order.placed" }

// OrderPaid is emitted when payment is recorded.
type OrderPaid struct {
	OrderID string
	Amount  money.Money
}

func (OrderPaid) EventName() string { return "order.paid" }

// OrderShipped is emitted when the order ships.
type OrderShipped struct {
	OrderID string
}

func (OrderShipped) EventName() string { return "order.shipped" }
