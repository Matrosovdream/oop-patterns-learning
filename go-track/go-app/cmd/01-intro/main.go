// Command 01-intro is the "does my environment work?" lesson.
//
// Run it:
//
//	docker compose run --rm golang go run ./cmd/01-intro
//	# or locally:
//	cd go-track/go-app && go run ./cmd/01-intro
package main

import (
	"fmt"

	"oop-patterns-learning/go-app/internal/money"
)

func main() {
	fmt.Println("Go track is alive. Welcome to OOP-by-design.")
	fmt.Println()

	// A first taste of the running example: an amount of money.
	price := money.FromMajor(19, 99, "USD")
	withTax := price.Percentage(108) // 108% == price + 8% tax
	fmt.Printf("list price:    %s\n", price)
	fmt.Printf("with 8%% tax:   %s\n", withTax)

	fmt.Println()
	fmt.Println("Next: lesson 02 — structs, methods & encapsulation.")
}
