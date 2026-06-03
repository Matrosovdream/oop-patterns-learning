// Command 08-dip-di shows dependency inversion + injection wired by hand at a
// "composition root" — the one place allowed to know concrete types.
package main

import (
	"fmt"
	"time"

	"oop-patterns-learning/go-app/internal/catalog"
	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/notify"
)

func main() {
	// ---- Composition root ----
	// main() is the only place that picks concrete implementations. It builds the
	// low-level details and injects them into the high-level policy (the Service).
	// Swap any line below for a different implementation and Service is unaffected.
	var repo catalog.Repository = catalog.NewInMemoryRepository()
	var sender notify.Sender = notify.EmailSender{From: "catalog@shopkit.test"}
	clock := func() time.Time { return time.Now() }

	svc := catalog.NewService(repo, sender, clock)

	// ---- Use the wired application ----
	if _, err := svc.AddProduct("MUG-1", "Enamel Mug", money.FromMajor(12, 0, "USD")); err != nil {
		panic(err)
	}
	if _, err := svc.AddProduct("TEE-1", "Logo Tee", money.FromMajor(25, 0, "USD")); err != nil {
		panic(err)
	}

	fmt.Printf("\ncatalog now has %d products.\n", svc.Count())
	fmt.Println("Service never named a database or a mailer — only abstractions.")
}
