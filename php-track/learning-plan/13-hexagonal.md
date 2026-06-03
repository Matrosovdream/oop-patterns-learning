# 13 · Ports & Adapters (Hexagonal)

## Goal
Make the boundary between your application and the outside world explicit with **ports**
(interfaces the core owns) and **adapters** (implementations that plug in) — and let Laravel's
container swap them.

## Theory
Hexagonal architecture (Alistair Cockburn, "Ports & Adapters") pictures your application as a
**hexagon**. Inside is business logic. The edges are **ports** — interfaces the core defines in
its own language. Outside, **adapters** implement those ports to talk to specific technologies.

Two kinds of port:
- **Driving (inbound)** — how the outside calls *in* (a use case; an HTTP controller is a
  driving adapter that calls it).
- **Driven (outbound)** — what the core needs *from* the outside (`OrderRepository`,
  `EventPublisher`). The core declares the interface; a database/queue adapter implements it.

The payoff is **dependency inversion at system scale** (lesson 08, grown up): the core depends
only on its ports, so you run it against in-memory adapters in tests and real ones in prod, and
swap technologies (in-memory → Eloquent, log → queue) without touching business logic.

## Language notes (PHP / Laravel)
- Ports are **interfaces** (`App\Ordering\Domain\OrderRepository`, `App\Ordering\Application\
  EventPublisher`); adapters are classes that `implements` them
  (`InMemoryOrderRepository`, `LoggingPublisher`, `RecordingPublisher`).
- The **service container** binds a port to an adapter in one line:
  `$this->app->singleton(OrderRepository::class, InMemoryOrderRepository::class);` — swapping
  adapters is editing that line, not the use cases.
- A **test double is just another adapter**: `RecordingPublisher` captures events instead of
  logging them. No mocking library needed; bind it and assert on `names()`.
- `FanOutPublisher` is itself an adapter that forwards to several others (Composite, lesson 10)
  — so one event stream feeds log + recorder + read model.

## Practice
```bash
docker compose run --rm laravel php artisan ordering:demo
```

Read `app/Providers/OrderingServiceProvider.php` (the bindings), the two ports
(`Domain/OrderRepository.php`, `Application/EventPublisher.php`), and their adapters under
`Infrastructure/`. Notice the use cases (`PlaceOrderHandler`, …) name only the ports.

## Exercises
1. Write an `EloquentOrderRepository implements OrderRepository` (it can wrap a model or even a
   JSON column) and change the one binding line. The handlers don't change.
2. In a quick test/tinker, bind `EventPublisher` to *only* a `RecordingPublisher`, place an
   order, and assert the captured event names. You just tested a use case with no framework
   mocking.
3. Label every port in this codebase as driving or driven.

## Recap & next
Explicit ports, pluggable adapters, container-wired. Next we generalize this into the
dependency rule of Clean/Onion architecture.

← [12 · Layered](12-layered.md) · → [14 · Clean / Onion architecture](14-clean.md)
