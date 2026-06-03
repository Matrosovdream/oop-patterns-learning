# 14 · Clean / Onion architecture & the dependency rule

## Goal
State the one rule that unifies layered, hexagonal, onion, and "clean" architecture — the
**dependency rule** — and recognize when the full ceremony is worth it.

## Theory
Onion (Jeffrey Palermo) and Clean (Robert C. Martin) architecture are the same idea drawn as
concentric rings instead of a stack or a hexagon:

```
        ┌─────────────────────────────┐
        │  frameworks / drivers (UI,   │   outer: volatile detail
        │   DB, web) — adapters        │
        │   ┌─────────────────────┐    │
        │   │  interface adapters  │   │
        │   │   ┌──────────────┐   │   │
        │   │   │ use cases    │   │   │
        │   │   │  ┌────────┐  │   │   │
        │   │   │  │ entities│  │   │   │   inner: stable policy
        │   │   │  └────────┘  │   │   │
        │   │   └──────────────┘   │   │
        │   └─────────────────────┘    │
        └─────────────────────────────┘
```

**The dependency rule:** source dependencies point *only inward*. Inner rings know nothing
about outer ones. Entities (domain) don't know about use cases; use cases don't know about
the web or the database. When an inner ring needs something from outside, it defines an
*interface* (a port) and the outer ring implements it — control flows out, but the source
dependency still points in. This is the Dependency Inversion Principle applied as the
organizing law of the whole system.

What you buy: the core is **independent of frameworks, UI, and database** and therefore
**testable in isolation** and **durable** as technologies churn. What you pay: more
interfaces, DTOs, and mapping. Clean architecture is overkill for a CRUD admin panel and
invaluable for a complex, long-lived domain. Match the ceremony to the stakes.

## Language notes (Go)
- Go expresses "inner ring" with the import graph and `internal/`. The build *enforces* the
  dependency rule: an inward-pointing cycle won't compile.
- You can prove conformance at compile time with interface assertions:
  `var _ domain.Repository = (*infra.InMemoryOrders)(nil)` — if an outer type stops satisfying
  an inner port, the build breaks. The demo does exactly this.
- Idiomatic Go resists over-ringing. The community mantra (Three Dots Labs, etc.): take the
  *dependency rule*, skip the dogmatic four-ring folder taxonomy. domain / app / infra / cmd
  is usually plenty.

## Practice
```bash
docker compose run --rm golang go run ./cmd/14-clean
```

Read `cmd/14-clean/main.go`. The `var _ domain.Repository = (*infra.InMemoryOrders)(nil)`
lines are the dependency rule made executable: the outer adapter is checked against the inner
port at compile time. The run drives a full place→pay→ship lifecycle through use cases.

## Exercises
1. Break the rule deliberately: add an `import` of `infra` to a file in `app` that doesn't
   need it, and see whether it creates a cycle or just a smell. Revert.
2. Argue for or against full clean architecture for: (a) a 3-table internal CRUD tool, (b) a
   billing engine with 5 years of accreting rules. Use "rate of change" as your axis.
3. Add a compile-time assertion that your new `FileOrders` adapter (from lesson 13) satisfies
   `domain.Repository`.

## Recap & next
You can structure an application so its core outlives its frameworks. Part 5 fills that core
with real modeling power: **Domain-Driven Design**.

← [13 · Hexagonal](13-hexagonal.md) · → [15 · Strategic DDD](15-strategic-ddd.md)
