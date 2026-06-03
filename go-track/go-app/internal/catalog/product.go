// Package catalog holds the Product type from the ShopKit domain.
//
// Product demonstrates composition: it EMBEDS audit.Timestamps (gaining Touch,
// CreatedAt, UpdatedAt) and HOLDS a money.Money price as a field. Neither is
// inheritance — there is no base "Entity" class Product extends.
package catalog

import (
	"time"

	"oop-patterns-learning/go-app/internal/audit"
	"oop-patterns-learning/go-app/internal/money"
)

// Product is a catalog item.
type Product struct {
	audit.Timestamps // embedded: Product.Touch(), Product.CreatedAt, etc. are promoted

	SKU   string
	Name  string
	Price money.Money
}

// NewProduct constructs a Product and stamps its timestamps.
func NewProduct(sku, name string, price money.Money, now time.Time) *Product {
	p := &Product{SKU: sku, Name: name, Price: price}
	p.Touch(now) // promoted method from the embedded Timestamps
	return p
}

// Reprice changes the price and re-stamps UpdatedAt.
func (p *Product) Reprice(newPrice money.Money, now time.Time) {
	p.Price = newPrice
	p.Touch(now)
}
