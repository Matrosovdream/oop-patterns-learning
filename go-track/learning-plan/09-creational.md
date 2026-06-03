# 09 · Creational patterns

## Goal
Recognize the four GoF creational patterns — Factory Method, Abstract Factory, Builder,
Singleton — and, crucially, how little ceremony they need in Go.

## Theory
Creational patterns answer "how is this object made, and who decides the concrete type?"
The shared goal is to **decouple callers from concrete construction** so you can change what
gets built without changing who asks for it.

- **Factory Method** — a function/method returns an object of an interface type; the caller
  names *what* it wants, not *which class*. Add a new variant in one place.
- **Abstract Factory** — produce a *family* of related objects that must be used together
  (e.g. a payment charger + matching refunder). Prevents mixing incompatible parts.
- **Builder** — assemble a complex object step by step, often with optional parts and a
  fluent API. Separates "how to construct" from "the final representation."
- **Singleton** — guarantee exactly one instance with a global access point. The most abused
  pattern: a singleton is a global, and globals make testing and DI harder. Reach for it only
  for genuinely process-wide, immutable-ish things (and prefer injecting even those).

## Language notes (Go)
- **Factory Method is just a function** returning an interface — `NewPayment(kind) (Payment,
  error)`. There's no "creator class hierarchy"; the function is the pattern.
- **Abstract Factory is a struct/interface of constructors.** Each concrete factory returns a
  matching family.
- **Builder earns its place** when a type has many optional fields. (Go's other answer is the
  *functional options* idiom: `New(opts ...Option)` — look it up after this lesson.)
- **Singleton = `sync.Once` + a package var.** `sync.Once` makes lazy init concurrency-safe.
  But idiomatic Go often skips it: just construct the thing in `main` and inject it (lesson 08).

## Practice
```bash
docker compose run --rm golang go run ./cmd/09-creational
```

Read `internal/creational/creational.go` and `cmd/09-creational/main.go`. Note that the
concrete types (`cardPayment`, `stripeCharger`, …) are unexported — callers only ever touch
the factory functions and interfaces.

## Exercises
1. Add a `paypal` gateway family (its own charger + refunder) and select it via `NewGateway`.
   No caller of `Gateway` should change.
2. Rewrite the `Config` singleton as an injected dependency (pass `*Config` into a service).
   Which version is easier to test, and why?
3. Convert `CartBuilder` to the functional-options style: `cart.New("USD", cart.WithLine(...))`.
   Compare ergonomics.

## Recap & next
You can decouple "what to build" from "how it's built." Next: patterns for composing objects
into larger structures.

← [08 · DIP & DI](08-dip-di.md) · → [10 · Structural patterns](10-structural.md)
