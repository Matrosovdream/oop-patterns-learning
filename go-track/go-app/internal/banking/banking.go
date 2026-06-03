// Package banking is a teaching example for Liskov Substitution (LSP) and
// Interface Segregation (ISP).
//
// The trick: instead of one fat Account interface that every account type must
// implement (forcing a Vault to provide a Withdraw it can't honor — an LSP
// violation waiting to happen), we segregate capabilities into tiny interfaces.
// A consumer asks for exactly the capability it needs, and the compiler refuses
// to pass a type that lacks it.
package banking

import (
	"errors"

	"oop-patterns-learning/go-app/internal/money"
)

// ErrInsufficientFunds is returned when a withdrawal exceeds the balance.
var ErrInsufficientFunds = errors.New("banking: insufficient funds")

// Segregated interfaces (ISP): one capability each.
type (
	Depositable  interface{ Deposit(m money.Money) }
	Withdrawable interface {
		Withdraw(m money.Money) error
	}
)

// Checking supports both deposits and withdrawals, so it satisfies BOTH
// Depositable and Withdrawable — and it honors the Withdrawable contract fully
// (LSP): a withdrawal either succeeds or returns a documented error.
type Checking struct{ balance money.Money }

func NewChecking(opening money.Money) *Checking { return &Checking{balance: opening} }
func (c *Checking) Balance() money.Money        { return c.balance }
func (c *Checking) Deposit(m money.Money)       { c.balance, _ = c.balance.Add(m) }

func (c *Checking) Withdraw(m money.Money) error {
	if c.balance.Amount() < m.Amount() {
		return ErrInsufficientFunds
	}
	c.balance, _ = c.balance.Sub(m)
	return nil
}

// Vault can only RECEIVE money. It deliberately has no Withdraw method, so it
// satisfies Depositable but NOT Withdrawable. You therefore cannot pass it where
// a withdrawal is required — the mistake is caught at compile time, not as a
// surprise "operation not supported" panic at runtime.
type Vault struct{ balance money.Money }

// NewVault seeds the balance with a zero amount in the given currency, so the
// first Deposit doesn't hit a currency mismatch.
func NewVault(currency string) *Vault  { return &Vault{balance: money.New(0, currency)} }
func (v *Vault) Balance() money.Money  { return v.balance }
func (v *Vault) Deposit(m money.Money) { v.balance, _ = v.balance.Add(m) }
