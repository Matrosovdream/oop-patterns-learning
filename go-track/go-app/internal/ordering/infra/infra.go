// Package infra holds the ADAPTERS of the Ordering context: concrete
// implementations of the ports the inner layers defined. This is the outermost
// layer — it imports domain and app, never the other way around. Swap an adapter
// (in-memory → SQL, logger → message bus) and the inner layers don't change.
package infra

import (
	"fmt"
	"sync"

	"oop-patterns-learning/go-app/internal/ordering/app"
	"oop-patterns-learning/go-app/internal/ordering/domain"
)

// InMemoryOrders is a driven adapter implementing the domain.Repository port
// with a map. A SQL or HTTP adapter would implement the same interface.
type InMemoryOrders struct {
	mu    sync.Mutex
	items map[domain.OrderID]*domain.Order
}

func NewInMemoryOrders() *InMemoryOrders {
	return &InMemoryOrders{items: map[domain.OrderID]*domain.Order{}}
}

func (r *InMemoryOrders) Save(o *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[o.ID()] = o
	return nil
}

func (r *InMemoryOrders) Find(id domain.OrderID) (*domain.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	o, ok := r.items[id]
	if !ok {
		return nil, domain.ErrOrderNotFound
	}
	return o, nil
}

// LoggingPublisher is a driven adapter for the app.EventPublisher port: it just
// prints events. A real one might push to Kafka, NATS, or an outbox table.
type LoggingPublisher struct{}

func (LoggingPublisher) Publish(events []domain.Event) {
	for _, e := range events {
		fmt.Printf("  event: %s\n", e.EventName())
	}
}

// FanOutPublisher forwards events to several publishers — so the command side
// can drive both a logger and the CQRS read-model projection at once. (It's the
// Composite pattern from lesson 10, applied to ports.)
type FanOutPublisher struct {
	Targets []app.EventPublisher
}

func (f FanOutPublisher) Publish(events []domain.Event) {
	for _, t := range f.Targets {
		t.Publish(events)
	}
}
