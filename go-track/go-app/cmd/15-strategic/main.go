// Command 15-strategic is a map, not a machine: it prints ShopKit's bounded
// contexts and how they relate. Strategic DDD is about drawing these boundaries
// BEFORE writing classes.
package main

import "fmt"

type boundedContext struct {
	name     string
	language string
	owns     string
}

func main() {
	contexts := []boundedContext{
		{"Catalog", "Product, SKU, Category", "what we sell"},
		{"Pricing", "Money, Discount, Tax", "what it costs"},
		{"Ordering", "Order, Line, OrderID, paid/shipped", "the purchase lifecycle"},
	}

	fmt.Println("ShopKit bounded contexts (each with its own ubiquitous language):")
	for _, c := range contexts {
		fmt.Printf("  %-9s | %-36s | %s\n", c.name, c.language, c.owns)
	}

	fmt.Println("\nContext map (the relationships between them):")
	fmt.Println("  Ordering --uses--> Pricing    Customer/Supplier: Ordering needs prices")
	fmt.Println("  Ordering --ACL---> Catalog    Anti-Corruption Layer translates products in")

	fmt.Println("\nKey insight: a 'Product' (Catalog) and a 'Line' (Ordering) are DIFFERENT")
	fmt.Println("models on purpose. The same real-world thing is shaped to each context's needs.")
}
