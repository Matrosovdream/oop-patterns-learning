// Command 04-composition shows Go's two composition tools and why there is no
// "inheritance" to miss.
package main

import (
	"fmt"
	"time"

	"oop-patterns-learning/go-app/internal/catalog"
	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/notify"
)

// PriceWatcher HOLDS a notify.Sender as a field (composition by dependency).
// It reuses sending behavior without being an EmailSender or SMSSender itself.
type PriceWatcher struct {
	sender notify.Sender
}

func (w PriceWatcher) AnnounceReprice(p *catalog.Product, to string) {
	msg := fmt.Sprintf("%s is now %s", p.Name, p.Price)
	_ = w.sender.Send(to, msg)
}

func main() {
	now := time.Now()

	// Product EMBEDS audit.Timestamps, so it has Touch/CreatedAt/UpdatedAt
	// "for free" — composition by embedding.
	mug := catalog.NewProduct("MUG-1", "Enamel Mug", money.FromMajor(12, 00, "USD"), now)
	fmt.Printf("created %s at %s\n", mug.Name, mug.CreatedAt.Format(time.Kitchen))

	mug.Reprice(money.FromMajor(9, 50, "USD"), now.Add(time.Hour))
	fmt.Printf("repriced; updated at %s\n", mug.UpdatedAt.Format(time.Kitchen))

	// PriceWatcher composes an emailer in by field — swap it for SMS freely.
	watcher := PriceWatcher{sender: notify.EmailSender{From: "deals@shopkit.test"}}
	watcher.AnnounceReprice(mug, "fan@example.com")

	fmt.Println("\nNo base classes, no `extends` — just embedding and fields.")
}
