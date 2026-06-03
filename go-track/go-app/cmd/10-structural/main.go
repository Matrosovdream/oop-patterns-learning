// Command 10-structural exercises Adapter, Decorator, Composite, Proxy, Facade.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/cart"
	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/notify"
	"oop-patterns-learning/go-app/internal/pricing"
	"oop-patterns-learning/go-app/internal/structural"
)

func main() {
	// Adapter: a legacy SMS client now behaves like a notify.Sender.
	var sms notify.Sender = structural.SMSAdapter{Client: structural.LegacySMS{}}

	// Decorator: wrap senders to add a tag — same interface, extra behavior.
	tagged := structural.PrefixSender{Inner: sms, Prefix: "[ShopKit] "}

	// Composite: email + (decorated) SMS treated as one Sender.
	group := structural.SenderGroup{
		notify.EmailSender{From: "shop@shopkit.test"},
		tagged,
	}
	fmt.Println("adapter + decorator + composite:")
	_ = group.Send("buyer@example.com", "Your order shipped!")

	// Proxy: the caching proxy hits the slow backend only once per currency.
	backend := &structural.SlowRates{}
	rates := structural.NewCachingRates(backend)
	for i := 0; i < 3; i++ {
		_ = rates.Rate("USD")
	}
	fmt.Printf("\nproxy: 3 lookups, backend called %d time(s)\n", backend.Calls)

	// Facade: one call hides the cart→discount→notify orchestration.
	cartObj := cart.New("USD")
	cartObj.Add(cart.Line{SKU: "MUG-1", Name: "Enamel Mug", Unit: money.FromMajor(12, 0, "USD"), Quantity: 2})
	c := structural.Checkout{Notifier: notify.EmailSender{From: "orders@shopkit.test"}}
	total := c.PlaceOrder(cartObj, pricing.PercentageOff{Percent: 10}, "buyer@example.com")
	fmt.Printf("\nfacade: charged %s\n", total)
}
