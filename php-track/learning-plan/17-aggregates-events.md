# 17 · Aggregates, repositories & domain events

## Goal
Group entities and value objects into an **aggregate** with a consistency boundary, load and
store it whole through a **repository**, and announce change with **domain events**.

## Theory
**Aggregate.** A cluster of objects treated as one unit for changes, with one entity as the
**aggregate root**. The root is the only entry point: outside code holds the root and calls
*its* methods; it never reaches in to mutate a child directly. That's how you enforce
**invariants** — rules true for the whole cluster ("the total equals the sum of lines", "you
can't add lines to a placed order"). `Order` is the root; `Line`s live inside it; you change
them only via `Order::addLine()`.

Design rules of thumb:
- **Keep aggregates small.** The boundary is a *consistency* boundary, not "everything
  related." Big aggregates mean lock contention and rigid loading.
- **Reference other aggregates by id**, not by object. An order references a customer by
  `CustomerId`, not an embedded `Customer`. One transaction changes one aggregate.

**Repository.** A collection-like interface for loading/saving *whole aggregates* by identity
(`save`, `find`). It speaks the domain's language, hides storage, and is a **port** defined in
the domain (lesson 13). In Laravel you resist the temptation to let the aggregate *be* an
Eloquent model — instead an Eloquent adapter maps rows ↔ aggregate.

**Domain event.** A record that something meaningful happened, past-tense, in the ubiquitous
language (`OrderPlaced`, `OrderPaid`). The aggregate *records* events as it changes; the
application *pulls and publishes* them after saving. Events decouple "what happened" from "who
cares" — the seam for side effects, read models (lesson 19), and cross-context integration.

## Language notes (PHP / Laravel)
- The aggregate root is a `final` class with **private state** and methods that are the only way
  to change it — exactly `App\Ordering\Domain\Order`. Private state is what *guarantees* the
  invariants can't be bypassed.
- Events are small `final` classes implementing a tiny `DomainEvent` interface; the root
  accumulates them and exposes `pullEvents()` (return + clear).
- The repository is an **interface in `Domain`**; the adapter lives in `Infrastructure`. Don't
  conflate it with Laravel's Eloquent: Eloquent is a great *adapter*, a poor *domain model*.
- `lines()` returns the internal array; in PHP arrays are copied by value on return, so callers
  can't mutate the aggregate's storage — a free defensive win (objects inside are still shared,
  but `Line` is immutable).

## Practice
```bash
docker compose run --rm laravel php artisan ordering:demo
```

Read `app/Ordering/Domain/Order.php` (root + invariants + events), `Events/*`, and
`OrderRepository.php`. The demo's "ORD-3 (no lines)" line shows an invariant rejecting an
illegal state; the captured event list shows `order.placed/paid/shipped` recorded by the
aggregate.

## Exercises
1. Add an invariant: an order may hold at most 100 total units. Enforce it in `addLine()` and
   prove it can't be bypassed.
2. Make `Order` reference a customer by a `CustomerId` value object, not a `Customer` object.
   Why does "by id" matter for transactions?
3. Add a `LineAdded` event recorded by `addLine()`. Should it publish before or after `place()`?
   Why?

## Recap & next
Aggregates protect invariants; repositories persist them whole; events announce change. Next:
the application layer that drives them, and keeping foreign models out.

← [16 · Entities & value objects](16-entities-value-objects.md) · → [18 · Application layer & use cases](18-application-layer.md)
