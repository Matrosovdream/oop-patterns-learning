# 19 · CQRS & an introduction to event sourcing

## Goal
Separate the model that *changes* state from the model that *reads* it (CQRS), build a read
model by projecting domain events, and understand what event sourcing adds (and costs).

## Theory
**CQRS — Command Query Responsibility Segregation.** The write side and the read side have
different needs, so give them different models:

- **Command side** — changes state through aggregates, enforcing every invariant. Optimized
  for *correctness and consistency*. (Everything from lessons 16–18.)
- **Query side** — answers questions for display. Optimized for *reads*: denormalized, flat,
  often pre-joined. It holds no rules and never mutates aggregates.

A natural way to keep the read side fresh is **projection**: subscribe to the command side's
**domain events** and update read models as events arrive. The order book in the demo is built
purely by projecting `OrderPlaced` / `OrderPaid` / `OrderShipped`.

CQRS is a *spectrum*, not a switch. The light version (one database, separate read queries) is
cheap and common. The heavy version (separate read store, async projections, eventual
consistency) is powerful but adds real complexity — only worth it when read and write loads or
shapes genuinely diverge. Don't CQRS a CRUD app.

**Event sourcing** goes further: instead of storing current state, store the **sequence of
events** as the source of truth and *rebuild* state by replaying them. You gain a perfect
audit log, time-travel, and trivial projections; you pay with schema-evolution headaches,
snapshotting, and a steeper mental model. CQRS and event sourcing are *often* paired but are
independent choices — you can have either without the other.

## Language notes (Go)
- The read model is plain structs + a store (`OrderSummary`, `SummaryProjection`). The
  projection satisfies the same `app.EventPublisher` port as the logger, so you plug it in
  with a `FanOutPublisher` (the Composite pattern again) — one event stream, many consumers.
- Map iteration order is random in Go, so the query side sorts (`All()` sorts by ID) for stable
  output — a small reminder that read models exist to serve queries conveniently.
- Event sourcing's replay is a **fold over events** — `replay()` in the demo is literally a
  loop accumulating state. That simplicity is the whole appeal.

## Practice
```bash
docker compose run --rm golang go run ./cmd/19-cqrs
```

Read `internal/ordering/app/queries.go` and `cmd/19-cqrs/main.go`:
- Commands drive aggregates (write side); the `SummaryProjection` builds the read model from
  the emitted events.
- `readModel.All()` answers a query without ever touching an `Order` aggregate.
- `replay()` reconstructs status from an event stream — event sourcing in miniature.

## Exercises
1. Add a `revenue` total to the read model, updated on `OrderPaid`. Note you change only the
   projection — no aggregate, no command.
2. Add an `OrderCancelled` event and handle it in both the projection and `replay`. Did the
   write side need to know about the read side? (No — that decoupling is the point.)
3. List two systems where full event sourcing earns its complexity, and two where it would be
   a liability.

## Recap & next
Reads and writes can have different models, kept in sync by events. Time to put the entire
course together.

← [18 · Application layer](18-application-layer.md) · → [20 · Capstone](20-capstone.md)
