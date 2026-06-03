# 13 · Ports & Adapters (Hexagonal)

## Goal
Make the boundary between your application and the outside world explicit with **ports**
(interfaces the core owns) and **adapters** (implementations that plug in) — so the core is
testable and infrastructure is swappable.

## Theory
Hexagonal architecture (Alistair Cockburn, also "Ports & Adapters") reframes layering around
one image: your application is a **hexagon**. Inside is business logic. The edges are
**ports** — interfaces the core defines in its own language. Outside, **adapters** implement
those ports to talk to specific technologies.

Two kinds of port:
- **Driving (inbound) ports** — how the outside calls *in* (a use-case interface; an HTTP
  handler is a driving adapter that calls it).
- **Driven (outbound) ports** — what the core needs *from* the outside (a `Repository`, an
  `EventPublisher`). The core declares the interface; a database/queue adapter implements it.

The payoff is **dependency inversion at the system scale** (lesson 08, grown up): the core
depends only on its own ports, so you can run it against in-memory adapters in tests and real
ones in production, and swap technologies (Postgres → DynamoDB, email → SMS) without touching
business logic. The hexagon shape (vs a stack) is just a reminder that there's no privileged
"top" or "bottom" — every external thing is just another adapter on an edge.

## Language notes (Go)
- Go is almost purpose-built for this: **ports are interfaces, adapters are structs that
  satisfy them implicitly.** `domain.Repository` and `app.EventPublisher` are ports;
  `infra.InMemoryOrders` and `infra.LoggingPublisher` are adapters.
- Because interfaces are satisfied implicitly, an adapter never imports the port to "declare"
  it implements it — it just has the methods. The import arrow points from adapter → core.
- A **test double is just another adapter** (see the `recorder` in the demo). No mocking
  framework required.

## Practice
```bash
docker compose run --rm golang go run ./cmd/13-hexagonal
```

Read `cmd/13-hexagonal/main.go`: the same `PlaceOrderHandler` is driven once by the
`LoggingPublisher` adapter and once by a hand-written `recorder` adapter. The application code
is identical both times — only the adapter plugged into the port changed.

## Exercises
1. Write a `FileOrders` adapter that satisfies `domain.Repository` by writing JSON to a map
   keyed by ID (or a real file). Swap it in; the handlers don't change.
2. Identify the driving vs driven ports in this codebase. Which adapters drive the app, which
   are driven by it?
3. Sketch (on paper) the hexagon for ShopKit Ordering: core in the middle, ports on the
   edges, adapters outside.

## Recap & next
Explicit ports and pluggable adapters. Next we generalize this into the dependency rule that
defines Clean/Onion architecture.

← [12 · Layered](12-layered.md) · → [14 · Clean / Onion architecture](14-clean.md)
