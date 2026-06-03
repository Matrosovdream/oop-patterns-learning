# 15 · Strategic DDD

## Goal
Learn to carve a system into **bounded contexts** with their own **ubiquitous language**, and
to map the relationships between them — the strategic decisions you make *before* writing
tactical code.

## Theory
Domain-Driven Design has two halves. **Strategic** DDD (this lesson) is about the big picture:
where the boundaries go and what words mean inside them. **Tactical** DDD (lessons 16–19) is
the building blocks inside a boundary.

- **Ubiquitous language** — one shared vocabulary, used identically by domain experts, code,
  and conversation. If the business says "place an order", the method is `Place()`, not
  `submit()` or `createOrderRecord()`. Mismatched language is where bugs and miscommunication
  breed.
- **Bounded context** — an explicit boundary within which a model and its language are
  consistent. The same word means *different things* in different contexts: a "Product" in
  Catalog (name, description, images) is not the "Product" in Pricing (cost basis, tax class)
  or the "Line" in Ordering (sku, quantity, price-at-time-of-order). Trying to make one
  "Product" class serve all three is the classic path to an unmaintainable god-model.
- **Context map** — how contexts relate: *Shared Kernel*, *Customer/Supplier*, *Conformist*,
  *Anti-Corruption Layer* (ACL), *Open Host Service*, etc. The map is a political and
  technical document: it says who depends on whom and how each protects its model.

The deepest strategic move is **finding the boundaries** — usually along business
capabilities, not database tables. Get this wrong and no amount of clean tactical code saves
you.

## Language notes (Go)
- A bounded context maps cleanly to a **Go package tree** (`internal/ordering/...`,
  `internal/catalog/...`). The package boundary is a real, compiler-checked seam.
- Each context keeps its *own* types even when they model "the same" real-world thing — and
  translates at the edges (the ACL, lesson 18). Go's lack of a global object graph makes this
  separation natural; you only share what you explicitly import.
- For inter-context communication, **domain events** (lesson 17) crossing a package/service
  boundary keep contexts decoupled.

## Practice
```bash
docker compose run --rm golang go run ./cmd/15-strategic
```

Read `cmd/15-strategic/main.go`: it prints ShopKit's three contexts, each with its own
language, and a small context map. Compare the `Product` (Catalog) and `Line` (Ordering)
models in the code — deliberately different shapes for the same real thing.

## Exercises
1. Write the ubiquitous language glossary for the Ordering context: list every domain term
   (Order, Line, Place, Pay, Ship, …) and a one-line definition. Would a domain expert agree?
2. ShopKit adds "subscriptions". Is that a new bounded context or part of Ordering? Argue
   using language and rate-of-change.
3. Draw the context map and label each relationship (Customer/Supplier, ACL, …).

## Recap & next
You can split a system along meaning, not tables. Next: the tactical building blocks inside a
context, starting with the entity/value-object distinction.

← [14 · Clean architecture](14-clean.md) · → [16 · Entities & value objects](16-entities-value-objects.md)
