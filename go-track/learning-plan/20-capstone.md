# 20 · Capstone — a bounded context end to end

## Goal
See every idea in the course working together in one runnable system, and have a template you
can copy for a real bounded context.

## Theory
The capstone isn't new theory — it's the payoff. The Ordering context demonstrates, in one
flow:

- **OOP foundations** — `Money`/`Line` are encapsulated, immutable value objects;
  polymorphism via interfaces throughout (Parts 1).
- **SOLID** — single-responsibility packages; the handlers are open/closed over events;
  everything depends on abstractions injected at the composition root (Part 2).
- **Patterns** — Strategy (events → projections), Composite (`FanOutPublisher`), Factory
  (constructors), Command (use-case handlers) (Part 3).
- **Architecture** — domain / app / infra / cmd layering, ports & adapters, the dependency
  rule enforced by the import graph and `internal/` (Part 4).
- **DDD + CQRS** — a bounded context with its own language, an `Order` aggregate guarding
  invariants, a repository port, domain events, an anti-corruption layer, and a read model
  projected from events (Part 5).

The shape to internalize: **a small, rule-rich domain at the center; thin use cases around it;
adapters at the edge; events carrying news outward.** That shape scales from this demo to a
real service.

## Language notes (Go)
- The whole context lives under `internal/ordering/{domain,app,infra}` with `cmd/` entry
  points — a layout you can lift directly into a production Go service.
- Notice how little "framework" there is: interfaces, structs, functions, and `main` doing the
  wiring. That's idiomatic Go architecture — the patterns are present, the ceremony isn't.

## Practice
```bash
docker compose run --rm golang go run ./cmd/20-capstone
```

Read `cmd/20-capstone/main.go`. It places an order through its full lifecycle, leaves a second
pending, has a third correctly rejected by a domain invariant, and finally queries the CQRS
read model — all wired at one composition root.

## Your capstone challenge
Build a **new bounded context** of your choice (e.g. *Shipping*, *Inventory*, or
*Subscriptions*) using this one as a template:

1. **Strategically**: name the context, write its ubiquitous language glossary, and place it
   on the ShopKit context map.
2. **Tactically**: design the aggregate root and its invariants; identify the value objects;
   list the domain events.
3. **Structurally**: lay out `domain` / `app` / `infra` packages; define the repository port;
   write at least two use cases.
4. **Integrate**: consume Ordering's events (e.g. *Shipping* reacts to `OrderPaid`) through an
   anti-corruption translation — don't import Ordering's domain into yours.
5. **Read side**: add one projection/read model and a query.

Run it from a new `cmd/`. When the import graph only points inward and your aggregate refuses
every illegal move, you've internalized the course.

## Recap
You started with structs and methods and finished with a clean, event-driven, domain-centric
bounded context — in a language with no classes and no inheritance. The design ideas were the
real subject; Go just kept them honest. Go build something.

← [19 · CQRS & event sourcing](19-cqrs.md) · [Back to the Go plan](README.md)
