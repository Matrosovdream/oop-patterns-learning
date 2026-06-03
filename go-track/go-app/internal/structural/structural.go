// Package structural shows the GoF structural patterns. Most are about wiring
// objects together behind a stable interface — and interfaces are Go's native
// idiom, so these read very naturally here.
package structural

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/cart"
	"oop-patterns-learning/go-app/internal/money"
	"oop-patterns-learning/go-app/internal/notify"
	"oop-patterns-learning/go-app/internal/pricing"
)

// ---------------------------------------------------------------------------
// Adapter — make an incompatible type fit an interface you already use.
// ---------------------------------------------------------------------------

// LegacySMS is a pretend third-party client we don't control. Its method shape
// (SendText(number, body)) doesn't match our notify.Sender (Send(to, message)).
type LegacySMS struct{}

func (LegacySMS) SendText(number, body string) {
	fmt.Printf("  [legacy-sms] to %s: %q\n", number, body)
}

// SMSAdapter adapts LegacySMS to notify.Sender, so legacy code drops straight
// into anything that expects a Sender (Broadcast, decorators, composites…).
type SMSAdapter struct{ Client LegacySMS }

func (a SMSAdapter) Send(to, message string) error {
	a.Client.SendText(to, message)
	return nil
}

// ---------------------------------------------------------------------------
// Decorator — wrap a value to add behavior, keeping the SAME interface.
// ---------------------------------------------------------------------------

// PrefixSender wraps any Sender and tags every message. Because it is itself a
// Sender, you can stack decorators (prefix(retry(log(real)))).
type PrefixSender struct {
	Inner  notify.Sender
	Prefix string
}

func (d PrefixSender) Send(to, message string) error {
	return d.Inner.Send(to, d.Prefix+message)
}

// ---------------------------------------------------------------------------
// Composite — treat a group of objects like a single object.
// ---------------------------------------------------------------------------

// SenderGroup is a slice of Senders that is itself a Sender. Send fans out to
// every child, so callers can't tell one sender from many.
type SenderGroup []notify.Sender

func (g SenderGroup) Send(to, message string) error {
	for _, s := range g {
		if err := s.Send(to, message); err != nil {
			return err
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// Proxy — same interface as the real thing, but it controls access (here: a
// cache so the expensive backend is hit at most once per key).
// ---------------------------------------------------------------------------

// RateProvider looks up a tax rate (as a percentage). Pretend it's a slow API.
type RateProvider interface{ Rate(currency string) int64 }

// SlowRates simulates an expensive backend and counts how often it's called.
type SlowRates struct{ Calls int }

func (s *SlowRates) Rate(currency string) int64 {
	s.Calls++
	return 108 // 100% + 8% tax, for demo purposes
}

// CachingRates is a proxy in front of any RateProvider.
type CachingRates struct {
	Inner RateProvider
	cache map[string]int64
}

func NewCachingRates(inner RateProvider) *CachingRates {
	return &CachingRates{Inner: inner, cache: map[string]int64{}}
}

func (c *CachingRates) Rate(currency string) int64 {
	if v, ok := c.cache[currency]; ok {
		return v
	}
	v := c.Inner.Rate(currency)
	c.cache[currency] = v
	return v
}

// ---------------------------------------------------------------------------
// Facade — one simple entry point over a tangle of subsystems.
// ---------------------------------------------------------------------------

// Checkout is a facade over the cart, pricing, and notification subsystems.
// Callers get one method instead of orchestrating three packages themselves.
type Checkout struct {
	Notifier notify.Sender
}

func (co Checkout) PlaceOrder(c *cart.Cart, d pricing.Discount, email string) money.Money {
	total := d.Apply(c.Subtotal())
	_ = co.Notifier.Send(email, "Thanks! Your order total is "+total.String())
	return total
}
