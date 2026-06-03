# 18 · Application layer & use cases

## Goal
Write a thin application layer that *orchestrates* the domain through use cases and DTOs, and
protect the domain from foreign models with an **anti-corruption layer**.

## Theory
The **application layer** (a.k.a. use cases / interactors / command handlers) is the
choreographer. A use case does a fixed dance:

1. accept a **command/DTO** (plain input data),
2. load the aggregate(s) from a repository,
3. call **domain methods** (where the rules live),
4. save,
5. publish events.

Crucially, it holds **no business rules** — if you find an `if order.total > limit` in a
handler, that rule belongs in the domain. The application layer decides *transaction
boundaries* and *orchestration*; the domain decides *what's valid*.

**DTOs (data transfer objects)** are plain structs that cross the boundary in and out. They
are *not* domain types. Keeping them separate stops the outside world's shapes (JSON fields,
form names, another context's model) from leaking into and corrupting your model.

**Anti-Corruption Layer (ACL).** When you integrate with another context or a legacy system,
don't let its model into yours. Put a translation layer at the boundary that converts foreign
shapes into your DTOs/domain types. The Ordering context turns `catalog.Product` into its own
`PlaceOrderLine` — Catalog could rename every field tomorrow and only the translator changes.

## Language notes (Go)
- A use case is a small struct holding its ports (`PlaceOrderHandler` holds a `Repository` and
  an `EventPublisher`) with a single `Handle(cmd) (result, error)` method. No framework, no
  base class.
- Commands/DTOs are plain structs (`PlaceOrderCommand`, `PlaceOrderLine`). They carry
  primitives (`UnitMinor int64`), not domain value objects — the handler builds the value
  objects, which is itself a tiny ACL inward.
- The ACL is just a translation function (`translate(products) []app.PlaceOrderLine`). The
  Ordering domain never imports `catalog` — verify that in the import list.

## Practice
```bash
docker compose run --rm golang go run ./cmd/18-application-layer
```

Read `internal/ordering/app/commands.go` and `cmd/18-application-layer/main.go`:
- `PlaceOrderHandler.Handle` is pure choreography — load/build, call domain, save, publish.
- `translate(...)` is the ACL converting Catalog's model into Ordering's DTO.

## Exercises
1. Add a `CancelOrder` use case. Put the *rule* ("only pending/paid orders can cancel") in the
   `Order` aggregate and keep the handler thin. Notice the split.
2. Sneak a business rule into a handler (e.g. reject orders over $1000 there), then move it
   into the domain. Which version would a second use case reuse correctly?
3. The ACL currently assumes USD. Make `translate` carry each product's real currency and
   reject a mixed-currency order. Where does that rejection belong?

## Recap & next
Thin orchestration, fat domain, foreign models kept at the door. Next: split reads from writes
and let events build read models — CQRS.

← [17 · Aggregates & events](17-aggregates-events.md) · → [19 · CQRS & event sourcing](19-cqrs.md)
