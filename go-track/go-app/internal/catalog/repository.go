package catalog

// Repository is the abstraction the domain depends on for persistence. Note it's
// declared here, in terms the domain cares about (Save/Find/All a *Product), not
// in terms of SQL or HTTP. Concrete stores (in-memory now, a database later)
// implement it. This is the seam that lets lesson 13's hexagonal architecture
// swap an adapter without the core noticing.
type Repository interface {
	Save(p *Product) error
	Find(sku string) (*Product, bool)
	All() []*Product
}

// InMemoryRepository is a concrete Repository backed by a map. It keeps insertion
// order so All() is deterministic (handy for demos and tests).
type InMemoryRepository struct {
	items map[string]*Product
	order []string
}

// NewInMemoryRepository returns an empty in-memory store.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{items: map[string]*Product{}}
}

func (r *InMemoryRepository) Save(p *Product) error {
	if _, exists := r.items[p.SKU]; !exists {
		r.order = append(r.order, p.SKU)
	}
	r.items[p.SKU] = p
	return nil
}

func (r *InMemoryRepository) Find(sku string) (*Product, bool) {
	p, ok := r.items[sku]
	return p, ok
}

func (r *InMemoryRepository) All() []*Product {
	out := make([]*Product, 0, len(r.order))
	for _, sku := range r.order {
		out = append(out, r.items[sku])
	}
	return out
}
