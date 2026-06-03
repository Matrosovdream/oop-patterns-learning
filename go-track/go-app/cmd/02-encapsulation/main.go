// Command 02-encapsulation shows why hiding state behind a constructor and
// methods matters, using the money.Money value type.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/money"
)

func main() {
	// We can only build Money through the constructor — the fields are
	// unexported, so no caller can put it into an invalid state from outside.
	wallet := money.FromMajor(50, 00, "USD")
	coffee := money.FromMajor(4, 75, "USD")

	remaining, err := wallet.Sub(coffee)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wallet %s - coffee %s = %s\n", wallet, coffee, remaining)

	// `wallet` is unchanged: Money is immutable, so Sub returned a NEW value.
	fmt.Printf("wallet is still %s (immutability)\n", wallet)

	// Encapsulation also lets the type *reject* nonsense. Mixing currencies
	// is a domain error, surfaced as a normal Go error value.
	euros := money.FromMajor(10, 00, "EUR")
	if _, err := wallet.Add(euros); err != nil {
		fmt.Printf("blocked bad operation: %v\n", err)
	}
}
