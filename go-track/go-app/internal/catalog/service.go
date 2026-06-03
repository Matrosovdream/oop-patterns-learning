package catalog

import (
	"fmt"
	"time"

	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/notify"
)

// Clock abstracts "what time is it" so the service never calls time.Now directly.
// Inject a fixed clock and behavior becomes deterministic.
type Clock func() time.Time

// Service is HIGH-LEVEL policy. The Dependency Inversion Principle says high-level
// policy must not depend on low-level details — both should depend on
// abstractions. So Service depends only on the Repository interface, the
// notify.Sender interface, and a Clock. It names no database, no mailer, and not
// even the wall clock. The concrete choices are injected from the outside.
type Service struct {
	repo   Repository
	sender notify.Sender
	now    Clock
}

// NewService injects the collaborators (constructor injection).
func NewService(repo Repository, sender notify.Sender, now Clock) *Service {
	return &Service{repo: repo, sender: sender, now: now}
}

// AddProduct creates, stores, and announces a product.
func (s *Service) AddProduct(sku, name string, price money.Money) (*Product, error) {
	p := NewProduct(sku, name, price, s.now())
	if err := s.repo.Save(p); err != nil {
		return nil, fmt.Errorf("save product %s: %w", sku, err)
	}
	_ = s.sender.Send("ops@shopkit.test", fmt.Sprintf("new product: %s @ %s", p.Name, p.Price))
	return p, nil
}

// Count reports how many products exist (delegated to the repository).
func (s *Service) Count() int { return len(s.repo.All()) }
