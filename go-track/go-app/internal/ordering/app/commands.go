// Package app is the application layer of the Ordering context. It holds USE
// CASES (command handlers), orchestrating the domain and the ports. It contains
// no business rules — those live in the domain — only the choreography: load,
// call a domain method, save, publish.
//
// It depends inward on the domain, and defines its own output ports
// (EventPublisher) that infrastructure adapters implement. It never imports
// infrastructure.
package app

import (
	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/ordering/domain"
)

// EventPublisher is an OUTPUT PORT. The application hands domain events to "the
// outside" without knowing whether they're logged, queued, or projected into a
// read model. Adapters (lesson 13) implement it.
type EventPublisher interface {
	Publish(events []domain.Event)
}

// ---- PlaceOrder use case ----

// PlaceOrderLine is an input DTO — plain transport data, not a domain type.
// Keeping DTOs separate from domain objects is the anti-corruption boundary
// (lesson 18): the outside world's shapes don't leak into the model.
type PlaceOrderLine struct {
	SKU       string
	Name      string
	UnitMinor int64 // price in minor units (cents)
	Quantity  int64
}

// PlaceOrderCommand is the intent "place this order".
type PlaceOrderCommand struct {
	OrderID  string
	Currency string
	Lines    []PlaceOrderLine
}

// PlaceOrderHandler executes the PlaceOrder use case.
type PlaceOrderHandler struct {
	orders    domain.Repository
	publisher EventPublisher
}

func NewPlaceOrderHandler(orders domain.Repository, pub EventPublisher) *PlaceOrderHandler {
	return &PlaceOrderHandler{orders: orders, publisher: pub}
}

func (h *PlaceOrderHandler) Handle(cmd PlaceOrderCommand) (domain.OrderID, error) {
	order := domain.NewOrder(domain.OrderID(cmd.OrderID), cmd.Currency)
	for _, l := range cmd.Lines {
		line := domain.Line{
			SKU:      l.SKU,
			Name:     l.Name,
			Unit:     money.New(l.UnitMinor, cmd.Currency),
			Quantity: l.Quantity,
		}
		if err := order.AddLine(line); err != nil {
			return "", err
		}
	}
	if err := order.Place(); err != nil {
		return "", err
	}
	if err := h.orders.Save(order); err != nil {
		return "", err
	}
	h.publisher.Publish(order.PullEvents())
	return order.ID(), nil
}

// ---- PayOrder use case ----

type PayOrderHandler struct {
	orders    domain.Repository
	publisher EventPublisher
}

func NewPayOrderHandler(orders domain.Repository, pub EventPublisher) *PayOrderHandler {
	return &PayOrderHandler{orders: orders, publisher: pub}
}

func (h *PayOrderHandler) Handle(orderID string) error {
	return mutate(h.orders, h.publisher, orderID, func(o *domain.Order) error { return o.Pay() })
}

// ---- ShipOrder use case ----

type ShipOrderHandler struct {
	orders    domain.Repository
	publisher EventPublisher
}

func NewShipOrderHandler(orders domain.Repository, pub EventPublisher) *ShipOrderHandler {
	return &ShipOrderHandler{orders: orders, publisher: pub}
}

func (h *ShipOrderHandler) Handle(orderID string) error {
	return mutate(h.orders, h.publisher, orderID, func(o *domain.Order) error { return o.Ship() })
}

// mutate is the shared load → change → save → publish choreography. The actual
// rule (Pay vs Ship) is the domain method passed in.
func mutate(orders domain.Repository, pub EventPublisher, id string, change func(*domain.Order) error) error {
	order, err := orders.Find(domain.OrderID(id))
	if err != nil {
		return err
	}
	if err := change(order); err != nil {
		return err
	}
	if err := orders.Save(order); err != nil {
		return err
	}
	pub.Publish(order.PullEvents())
	return nil
}
