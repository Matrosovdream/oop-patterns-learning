// Command 03-abstraction shows programming to an interface, not a concrete type.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/notify"
)

func main() {
	recipients := []string{"ada@example.com", "linus@example.com"}

	// Broadcast only knows about the notify.Sender abstraction. We hand it a
	// concrete EmailSender today…
	var sender notify.Sender = notify.EmailSender{From: "shop@shopkit.test"}
	fmt.Println("via email:")
	_ = notify.Broadcast(sender, recipients, "Your order shipped!")

	// …and swap in an SMSSender tomorrow with zero changes to Broadcast.
	sender = notify.SMSSender{Gateway: "twilio"}
	fmt.Println("via sms:")
	_ = notify.Broadcast(sender, recipients, "Your order shipped!")

	fmt.Println("\nSame Broadcast, two implementations — that's abstraction.")
}
