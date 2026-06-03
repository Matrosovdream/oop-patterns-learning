package app

import (
	"sort"
	"sync"

	"oop-patterns-learning/go-app/internal/ordering/domain"
)

// This file is the READ side of CQRS (lesson 19). Commands (commands.go) change
// state through the aggregate; queries read from a denormalized model optimized
// for display. The two sides never share a model.
//
// The read model here is built by PROJECTING the same domain events the command
// side emits. SummaryProjection satisfies the EventPublisher port, so you can
// plug it in right next to (or instead of) a logger.

// OrderSummary is a flat, display-ready view — no behavior, no invariants.
type OrderSummary struct {
	ID     string
	Status string
	Total  string
}

// SummaryProjection maintains OrderSummaries by reacting to domain events.
type SummaryProjection struct {
	mu        sync.Mutex
	summaries map[string]*OrderSummary
}

func NewSummaryProjection() *SummaryProjection {
	return &SummaryProjection{summaries: map[string]*OrderSummary{}}
}

// Publish makes the projection an EventPublisher: each event nudges the read
// model toward its new shape.
func (p *SummaryProjection) Publish(events []domain.Event) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, e := range events {
		switch ev := e.(type) {
		case domain.OrderPlaced:
			p.summaries[ev.OrderID] = &OrderSummary{ID: ev.OrderID, Status: string(domain.StatusPending), Total: ev.Total.String()}
		case domain.OrderPaid:
			if s := p.summaries[ev.OrderID]; s != nil {
				s.Status = string(domain.StatusPaid)
			}
		case domain.OrderShipped:
			if s := p.summaries[ev.OrderID]; s != nil {
				s.Status = string(domain.StatusShipped)
			}
		}
	}
}

// Get returns one summary.
func (p *SummaryProjection) Get(id string) (OrderSummary, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	s, ok := p.summaries[id]
	if !ok {
		return OrderSummary{}, false
	}
	return *s, true
}

// All returns every summary, ID-sorted for stable output.
func (p *SummaryProjection) All() []OrderSummary {
	p.mu.Lock()
	defer p.mu.Unlock()
	out := make([]OrderSummary, 0, len(p.summaries))
	for _, s := range p.summaries {
		out = append(out, *s)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
