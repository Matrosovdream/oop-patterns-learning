package domain

import "errors"

// ErrOrderNotFound is returned by a Repository when no order has the given ID.
var ErrOrderNotFound = errors.New("ordering: order not found")

// Repository is a PORT: an interface the domain defines in its own terms
// (save/find an Order by its identity), leaving the "how" (memory, SQL, an API)
// to an adapter in the infrastructure layer. Note the interface lives HERE, with
// the domain — the database package will depend on this, not the reverse. That
// inverted dependency is what hexagonal architecture (lesson 13) is built on.
//
// A repository deals in whole aggregates, never in rows or columns: you get back
// a fully-formed *Order with its invariants intact.
type Repository interface {
	Save(o *Order) error
	Find(id OrderID) (*Order, error)
}
