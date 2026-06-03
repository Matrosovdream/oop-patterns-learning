# 14 · Clean / Onion architecture & the dependency rule

## Goal
State the one rule that unifies layered, hexagonal, onion, and "clean" architecture — the
**dependency rule** — and use Laravel's container as your composition root.

## Theory
Onion (Jeffrey Palermo) and Clean (Robert C. Martin) architecture are the same idea drawn as
concentric rings:

```
  frameworks / drivers (Laravel, DB, HTTP) ── adapters   ← volatile detail (outer)
    interface adapters (controllers, repositories)
      use cases (application handlers)
        entities (domain model)                          ← stable policy (inner)
```

**The dependency rule:** source dependencies point *only inward*. Inner rings know nothing
about outer ones. The domain doesn't know about use cases; use cases don't know about HTTP or
the database. When an inner ring needs the outside, it defines an interface (a port) and the
outer ring implements it. This is the Dependency Inversion Principle (lesson 08) promoted to
the organizing law of the whole system.

What you buy: a core **independent of Laravel, HTTP, and the database** — testable in isolation
and durable as the framework churns. What you pay: more interfaces, DTOs, and mapping. Clean
architecture is overkill for a CRUD admin panel and invaluable for a complex, long-lived
domain. Match the ceremony to the stakes — and note most Laravel apps live happily *between*
"plain MVC" and "full clean", which is a fine place to be.

## Language notes (PHP / Laravel)
- Laravel's **service container is your composition root**. `OrderingServiceProvider::register()`
  is the *only* place that names concrete adapters; everything else type-hints interfaces and is
  auto-injected. That's DIP automated — you rarely write `new` for a dependency.
- Constructor injection + interface type-hints means the container resolves the whole graph for
  you (the controller asks for `PlaceOrderHandler`, which asks for `OrderRepository`, which the
  provider maps to an adapter).
- The dependency rule is **discipline, not compiler-enforced** in PHP (unlike Go's import
  cycles). Enforce it with a tool like Deptrac, or just code review: nothing under `Domain/`
  may `use Illuminate\*` or anything from `Infrastructure/`.

## Practice
```bash
docker compose run --rm laravel php artisan ordering:demo
```

Read `app/Providers/OrderingServiceProvider.php`. The bindings are the dependency rule made
real: inner-layer ports (`OrderRepository`, `EventPublisher`) mapped to outer-layer adapters.
The handlers and the controller never `new` an adapter — the container injects it.

## Exercises
1. Add a Deptrac (or hand-written) check that fails if `App\Ordering\Domain` imports
   `Illuminate` or `App\Ordering\Infrastructure`. Run it; keep it green.
2. Argue for/against full clean architecture for (a) a 3-table internal CRUD tool, (b) a billing
   engine with 5 years of accreting rules. Use "rate of change" as the axis.
3. Trace the injection graph the container builds when a request hits `/ordering`: controller →
   handler → repository/publisher → adapters. Which direction do the *source* dependencies point?

## Recap & next
You can structure a Laravel app so its core outlives the framework. Part 5 fills that core with
real modeling power: **Domain-Driven Design**.

← [13 · Hexagonal](13-hexagonal.md) · → [15 · Strategic DDD](15-strategic-ddd.md)
