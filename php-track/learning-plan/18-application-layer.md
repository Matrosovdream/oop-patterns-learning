# 18 · Application layer & use cases

## Goal
Write a thin application layer that *orchestrates* the domain through use cases and DTOs, and
protect the domain from foreign models with an **anti-corruption layer**.

## Theory
The **application layer** (use cases / command handlers) is the choreographer. A use case does
a fixed dance:

1. accept a **command/DTO** (plain input data),
2. load the aggregate(s) from a repository,
3. call **domain methods** (where the rules live),
4. save,
5. publish events.

It holds **no business rules** — find an `if ($order->total() > $limit)` in a handler and that
rule belongs in the domain. The application layer owns *transaction boundaries* and
*orchestration*; the domain owns *what's valid*.

**DTOs** are plain structs crossing the boundary in and out — *not* domain types. They stop the
outside world's shapes (JSON fields, form names, another context's model) from leaking into and
corrupting your model.

**Anti-Corruption Layer (ACL).** When integrating with another context or legacy system, don't
let its model into yours. Translate foreign shapes into your DTOs/domain types at the boundary.
Ordering turns a Catalog product into its own `PlaceOrderLine` — Catalog could rename every
field tomorrow and only the translator changes.

## Language notes (PHP / Laravel)
- A use case is a small class holding its ports (`PlaceOrderHandler` takes `OrderRepository` +
  `EventPublisher`) with one `handle($command)` method. The container injects the ports.
- Commands/DTOs are `final readonly`-style classes carrying primitives (`unitMinor: int`), not
  domain value objects — the handler builds the VOs, a tiny ACL inward.
- In Laravel, a **Form Request** validates and shapes input, then you map it to a command DTO —
  that mapping *is* an anti-corruption step at the HTTP edge. Keep the handler framework-free so
  it's reusable from HTTP, CLI, and queues alike (it already backs both our controller and
  Artisan command).
- Don't let a handler return an Eloquent model or a domain entity to the controller; return an
  id or a read DTO. The controller shapes the HTTP response.

## Practice
```bash
docker compose run --rm laravel php artisan ordering:demo
```

Read `app/Ordering/Application/PlaceOrderHandler.php` (pure choreography: build → call domain →
save → publish) and `OrderingController.php` / `OrderingDemo.php` (two presentations, same
handlers). For the ACL idea, see how the plain-PHP track's lesson 18 translated a `catalog
Product` into a DTO — the same translation belongs at Laravel's controller/Form-Request edge.

## Exercises
1. Add a `CancelOrder` use case. Put the *rule* ("only pending/paid orders can cancel") in the
   `Order` aggregate; keep the handler thin. Notice the split.
2. Add a Form Request that validates a place-order payload and maps it to `PlaceOrderCommand`.
   Where does validation stop and domain invariant-checking begin?
3. Sneak a business rule into the handler, then move it into the domain. Which version would a
   second use case reuse correctly?

## Recap & next
Thin orchestration, fat domain, foreign models kept at the door. Next: split reads from writes
and let events build read models — CQRS.

← [17 · Aggregates & events](17-aggregates-events.md) · → [19 · CQRS & event sourcing](19-cqrs.md)
