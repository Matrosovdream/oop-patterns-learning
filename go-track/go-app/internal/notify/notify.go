// Package notify demonstrates abstraction via a small, consumer-defined interface.
//
// The key idea: callers depend on the *behavior* "something I can send a message
// with", not on a concrete email/SMS type. In Go the interface is tiny and the
// concrete types never mention it — they satisfy it implicitly.
package notify

import "fmt"

// Sender is the abstraction: any type with this one method is a Sender.
// Small interfaces like this are the Go idiom ("the bigger the interface, the
// weaker the abstraction").
type Sender interface {
	Send(to, message string) error
}

// EmailSender is one concrete implementation. Note it does NOT say "implements
// Sender" anywhere — having the method is enough.
type EmailSender struct {
	From string
}

func (e EmailSender) Send(to, message string) error {
	fmt.Printf("  ✉  email from %s to %s: %q\n", e.From, to, message)
	return nil
}

// SMSSender is another implementation behind the same abstraction.
type SMSSender struct {
	Gateway string
}

func (s SMSSender) Send(to, message string) error {
	fmt.Printf("  📱 sms via %s to %s: %q\n", s.Gateway, to, message)
	return nil
}

// Broadcast works against the abstraction, so it can send through ANY Sender —
// today's email/SMS, tomorrow's Slack — without changing this code.
func Broadcast(s Sender, recipients []string, message string) error {
	for _, r := range recipients {
		if err := s.Send(r, message); err != nil {
			return fmt.Errorf("notify %s: %w", r, err)
		}
	}
	return nil
}
