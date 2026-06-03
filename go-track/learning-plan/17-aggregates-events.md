# 17 · Aggregates, repositories & domain events

## Goal
Group entities and value objects into an **aggregate** with a consistency boundary, load and
store it whole through a **repository**, and announce change with **domain events**.

## Theory
**Aggregate.** A cluster of objects treated as one unit for data changes, with one entity as
the **aggregate root**. The root is the only entry point: outside code holds a reference to
the root and calls *its* methods; it never reaches inside to mutate a child directly. This is
how you enforce **invariants** — rules that must hold for the whole cluster ("an order's total
equals the sum of its lines", "you can't add lines to a placed order"). `Order` is the root;
`Line`s live inside it; you change them only via `Order.AddLine`.

Design rules of thumb:
- **Keep aggregates small.** The boundary is a *consistency* boundary, not a "things that are
  related" boundary. Big aggregates mean lock contention and rigid loading.
- **Reference other aggregates by ID, not by pointer.** An order references a customer by
  `CustomerID`, it doesn't embed a `*Customer`. One transaction changes one aggregate.

**Repository.** A collection-like interface for loading/saving *whole aggregates* by identity
(`Save(o)`, `Find(id)`). It speaks the domain's language, hides the storage mechanism, and is
defined as a **port** in the domain (lesson 13). The domain depends on the interface; an
adapter implements it.

**Domain event.** A record that something meaningful happened, in the past tense and the
ubiquitous language (`OrderPlaced`, `OrderPaid`). The aggregate *records* events as it
changes; the application *pulls and publishes* them after saving. Events decouple "what
happened" from "who cares" — they're the seam for side effects, read models (lesson 19), and
cross-context integration (lesson 15).

## Language notes (Go)
- The aggregate root is a struct with **unexported state** and methods that are the only way
  to change it — exactly `domain.Order`. Unexported fields are what *guarantee* nothing
  bypasses the invariants.
- Events are small structs behind a tiny `Event` interface. The root accumulates them in an
  unexported slice and exposes `PullEvents()` (return + clear). Idiomatic and allocation-light.
- The repository is an interface in the domain package; the adapter (`infra.InMemoryOrders`)
  satisfies it implicitly. Note `Lines()` returns a *copy* so callers can't mutate internals —
  a small but important defensive habit in Go, where slices are reference-like.

## Practice
```bash
docker compose run --rm golang go run ./cmd/17-aggregates-events
```

Read `internal/ordering/domain/order.go` and `cmd/17-aggregates-events/main.go`. The demo
tries illegal moves (ship before pay, add a line after placing) and the aggregate refuses each
with a domain error, then walks the legal path and prints the events it recorded.

## Exercises
1. Add an invariant: an order may hold at most 100 total units. Enforce it in `AddLine` and
   prove it can't be bypassed from outside the package.
2. Make `Order` reference a customer by `CustomerID` (a value object), not a `*Customer`.
   Explain why "by ID" matters for transactions.
3. Add a `LineAdded` domain event recorded by `AddLine`. Should it be published before or
   after the order is placed? Why?

## Recap & next
Aggregates protect invariants; repositories persist them whole; events announce change. Next:
the application layer that drives them, and keeping foreign models out.

← [16 · Entities & value objects](16-entities-value-objects.md) · → [18 · Application layer & use cases](18-application-layer.md)
