# 20 · Capstone — a bounded context end to end

## Goal
See every idea in the course working together in one runnable Laravel feature, and have a
template you can copy for a real bounded context.

## Theory
The capstone isn't new theory — it's the payoff. The `App\Ordering` context demonstrates, in
one flow:

- **OOP foundations** — `Money`/`Line`/`OrderId` are encapsulated, immutable value objects;
  polymorphism via interfaces throughout (Part 1).
- **SOLID** — single-responsibility classes; handlers open/closed over events; everything
  depends on abstractions injected by the container (Part 2).
- **Patterns** — Factory (constructors), Composite (`FanOutPublisher`), Strategy/Observer
  (events → projections), Command (use-case handlers), State (the status lifecycle) (Part 3).
- **Architecture** — Domain / Application / Infrastructure layering, ports & adapters, the
  dependency rule, the container as composition root (Part 4).
- **DDD + CQRS** — a bounded context with its own language, an `Order` aggregate guarding
  invariants, a repository port, domain events, an anti-corruption boundary, and a read model
  projected from events (Part 5).

The shape to internalize: **a small, rule-rich domain at the center; thin use cases around it;
adapters at the edge; events carrying news outward.** It scales from this demo to a real
service — and it keeps Laravel at the edges instead of soaked through your business logic.

## Language notes (PHP / Laravel)
- The context lives under `app/Ordering/{Domain,Application,Infrastructure}`, wired by
  `OrderingServiceProvider`, and exposed through both an Artisan command and an HTTP route — a
  layout you can lift into a production Laravel app.
- Note how little Laravel leaks inward: only the Infrastructure adapters and the
  controller/command touch `Illuminate\*`. The domain is plain PHP.

## Practice
```bash
# CLI:
docker compose run --rm laravel php artisan ordering:demo
# HTTP:
docker compose up laravel        # then open http://localhost:8000/ordering
```

Read `app/Console/Commands/OrderingDemo.php` and `app/Http/Controllers/OrderingController.php`:
two presentations driving the *same* use cases. The CLI demo places an order through its full
lifecycle, leaves one pending, has one rejected by a domain invariant, then prints the CQRS read
model.

## Your capstone challenge
Build a **new bounded context** (e.g. *Shipping*, *Inventory*, or *Subscriptions*) using
Ordering as the template:

1. **Strategically**: name the context, write its ubiquitous-language glossary, place it on the
   ShopKit context map.
2. **Tactically**: design the aggregate root and its invariants; identify the value objects;
   list the domain events.
3. **Structurally**: lay out `Domain` / `Application` / `Infrastructure` namespaces; define the
   repository port; write at least two use cases.
4. **Wire it**: a service provider binding ports → adapters; expose a route or Artisan command.
5. **Integrate**: react to Ordering's events (e.g. *Shipping* listens for `OrderPaid`) through
   an anti-corruption translation — don't import Ordering's domain into yours.
6. **Read side**: add one projection/read model and a query.

When your domain has zero `Illuminate\*` imports and your aggregate refuses every illegal move,
you've internalized the course.

## Recap
You started with a `Money` class and finished with a clean, event-driven, domain-centric
bounded context in Laravel — with the framework kept firmly at the edges. The design ideas were
the real subject; Laravel is just the delivery mechanism. Go build something.

← [19 · CQRS & event sourcing](19-cqrs.md) · [Back to the PHP plan](README.md)
