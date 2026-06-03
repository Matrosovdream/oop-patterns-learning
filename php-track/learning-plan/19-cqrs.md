# 19 · CQRS & an introduction to event sourcing

## Goal
Separate the model that *changes* state from the model that *reads* it (CQRS), build a read
model by projecting domain events, and understand what event sourcing adds (and costs).

## Theory
**CQRS — Command Query Responsibility Segregation.** The write and read sides have different
needs, so give them different models:

- **Command side** — changes state through aggregates, enforcing every invariant. Optimized for
  *correctness*. (Everything from lessons 16–18.)
- **Query side** — answers questions for display. Optimized for *reads*: denormalized, flat,
  pre-joined. No rules, never mutates aggregates.

A natural way to keep the read side fresh is **projection**: subscribe to the command side's
**domain events** and update read models as they arrive. Our `SummaryProjection` builds the
order book purely from `OrderPlaced` / `OrderPaid` / `OrderShipped`.

CQRS is a *spectrum*. The light version (one database, separate read queries) is cheap and
common — and frankly where most Laravel apps should stop. The heavy version (separate read
store, async projections, eventual consistency) is powerful but costly; use it only when read
and write shapes/loads genuinely diverge. Don't CQRS a CRUD app.

**Event sourcing** goes further: store the **sequence of events** as the source of truth and
*rebuild* state by replaying them, instead of storing current state. You gain a perfect audit
log and time-travel; you pay with schema-evolution and snapshotting complexity. CQRS and event
sourcing are independent choices that often pair.

## Language notes (PHP / Laravel)
- The read model is plain classes (`OrderSummary`) + a store (`SummaryProjection`). The
  projection `implements EventPublisher`, so the container plugs it into the `FanOutPublisher`
  alongside the logger — one event stream, many consumers (Composite, lesson 10).
- Laravel's own **events + listeners** are a built-in observer mechanism; a listener that
  updates a read table is exactly a projection. (We hand-rolled it here to keep the pattern
  visible and framework-light.)
- Event sourcing's replay is a **fold over events** — a `match` in a loop accumulating state.
  For real event sourcing in Laravel, look at the `spatie/laravel-event-sourcing` package after
  this lesson.
- Query handlers should return read DTOs (`OrderSummary`), never aggregates.

## Practice
```bash
docker compose run --rm laravel php artisan ordering:demo
```

Read `app/Ordering/Application/SummaryProjection.php` and `OrderSummary.php`. In the demo, the
command side drives the aggregates while the projection builds the "order book" table you see —
and `OrderingController` answers the `/ordering` query straight from that read model, never
touching an `Order`.

## Exercises
1. Add a `revenue` total to `OrderSummary`, updated on `OrderPaid`. You change only the
   projection — no aggregate, no command.
2. Add an `OrderCancelled` event and handle it in the projection. Did the write side need to
   know about the read side? (No — that decoupling is the point.)
3. Write a tiny `replay(array $events): OrderStatus` that folds events into a status — event
   sourcing in miniature — then name two systems where full event sourcing earns its complexity.

## Recap & next
Reads and writes can have different models, kept in sync by events. Time to put the whole course
together.

← [18 · Application layer](18-application-layer.md) · → [20 · Capstone](20-capstone.md)
