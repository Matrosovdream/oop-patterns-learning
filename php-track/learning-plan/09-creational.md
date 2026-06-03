# 09 · Creational patterns

## Goal
Recognize and apply the four GoF creational patterns — Factory Method, Abstract Factory,
Builder, Singleton — and know when each is worth it in PHP.

## Theory
Creational patterns answer "how is this object made, and who decides the concrete type?" The
shared goal is to **decouple callers from concrete construction**, so you can change what gets
built without changing who asks.

- **Factory Method** — a method returns an object of an interface type; the caller names *what*
  it wants, not *which class*. New variant ⇒ one place to change.
- **Abstract Factory** — produce a *family* of related objects meant to be used together (a
  charger + matching refunder). Prevents mixing incompatible parts.
- **Builder** — assemble a complex object step by step, often fluent, with optional parts.
  Separates "how to construct" from "the finished representation."
- **Singleton** — exactly one instance with a global access point. The most-abused pattern: a
  singleton is a global that fights DI and testing. Reach for it rarely.

## Language notes (PHP)
- Factory Method / Abstract Factory are usually small classes with a `match` over a kind, or
  named static constructors on the product itself (`Money::fromMajor`).
- **Builder** shines when a constructor would have many optional args; the fluent `->add()->
  add()->build()` reads well. (Named arguments cover many simpler cases without a builder.)
- **Singleton**: a `private` constructor + `private static ?self $instance` + `instance()`. But
  in a framework, prefer a **container singleton binding** and inject it — you get "one
  instance" *without* the global. The lesson's `Config` notes this.
- Anonymous classes (`new class implements Charger {...}`) are handy for one-off family members
  (see `StripeGateway`).

## Practice
```bash
docker compose run --rm php php examples/09-creational.php
```

Read `src/Creational/*` and `examples/09-creational.php`. The concrete `CardPayment`/
`CashPayment` are only ever produced via `PaymentFactory` — callers depend on `Payment`.

## Exercises
1. Add a `paypal` gateway family (its own charger + refunder) selectable via `GatewayFactory`.
   No caller of `Gateway` should change.
2. Rewrite `Config` as a container-style injected dependency (pass it into a service). Which is
   easier to test, and why?
3. Replace `CartBuilder` with a constructor using named arguments. When is the builder still
   worth it?

## Recap & next
Decouple "what to build" from "how it's built." Next: patterns for composing objects into
larger structures.

← [08 · DIP & DI](08-dip-di.md) · → [10 · Structural patterns](10-structural.md)
