// Command 07-lsp-isp shows how segregated interfaces (ISP) make Liskov
// violations (LSP) impossible to write.
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/banking"
	"oop-patterns-learning/go-app/internal/money"
)

// transfer needs to take money OUT of one account and put it INTO another. It
// asks for the two narrow capabilities it actually uses — not a fat "Account".
// Anything Withdrawable+Depositable works; nothing else compiles.
func transfer(from banking.Withdrawable, to banking.Depositable, amount money.Money) error {
	if err := from.Withdraw(amount); err != nil {
		return err
	}
	to.Deposit(amount)
	return nil
}

func main() {
	checking := banking.NewChecking(money.FromMajor(100, 0, "USD"))
	vault := banking.NewVault("USD")

	// Checking is Withdrawable AND Depositable, so this is fine.
	if err := transfer(checking, vault, money.FromMajor(30, 0, "USD")); err != nil {
		fmt.Println("transfer failed:", err)
	}
	fmt.Printf("checking: %s, vault: %s\n", checking.Balance(), vault.Balance())

	// The LSP win: you literally cannot write `transfer(vault, checking, ...)`.
	// Vault has no Withdraw, so it does not satisfy Withdrawable, so it won't
	// compile. Uncomment to see the compiler reject it:
	//
	//   _ = transfer(vault, checking, money.FromMajor(10, 0, "USD"))
	//   // cannot use vault (*banking.Vault) as banking.Withdrawable value:
	//   // missing method Withdraw

	// Over-withdrawing returns a normal error (the Withdrawable contract honored).
	if err := transfer(checking, vault, money.FromMajor(1000, 0, "USD")); err != nil {
		fmt.Println("blocked over-withdrawal:", err)
	}
}
