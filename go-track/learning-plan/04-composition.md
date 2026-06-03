# 04 · Composition over inheritance

## Goal
Reuse behavior two ways — **embedding** and **holding a field** — and understand why Go
deliberately omits inheritance, the single biggest design lever in OOP.

## Theory
Classical OOP offers two ways to reuse code:

- **Inheritance** ("is-a"): `class SavingsAccount extends Account`. Powerful but rigid —
  subclasses are coupled to the base class's internals, deep hierarchies get fragile (the
  "fragile base class" problem), and you can only extend one parent.
- **Composition** ("has-a" / "uses-a"): an object holds other objects and delegates to them.

The industry learned the hard way to **favor composition over inheritance** (it's the first
principle in the GoF book). Composition is more flexible: you can swap the composed part at
runtime, combine many parts, and change one without disturbing the others. Inheritance
should be reserved for genuine, stable "is-a substitutable for" relationships — which are
rarer than they look (lesson 07's Liskov principle is the test).

Go takes a strong stance: **there is no inheritance.** You compose, full stop.

## Language notes (Go)
Two composition tools:

1. **Embedding** — put a type in a struct with no field name:
   ```go
   type Product struct {
       audit.Timestamps   // embedded
       SKU string
   }
   ```
   `Product` *promotes* `Timestamps`'s fields and methods: `p.Touch(now)`, `p.CreatedAt`
   work directly. This *looks* like inheritance but isn't — there's no subtype relationship,
   no overriding, and `Timestamps` knows nothing about `Product`. It's delegation with
   sugar.

2. **A plain field** — hold a dependency and call it:
   ```go
   type PriceWatcher struct { sender notify.Sender }
   ```
   `PriceWatcher` reuses sending behavior by *having* a `Sender`, not by *being* one. Swap
   the field for any other `Sender` and behavior changes with zero edits to `PriceWatcher`.

When you reach for embedding vs a field: embed when the outer type genuinely *is* an
extended version that wants the inner API on its surface; use a field when you just need to
*use* the dependency (the common case — and the more decoupled one).

## Practice
```bash
docker compose run --rm golang go run ./cmd/04-composition
```

Read `internal/catalog/product.go`, `internal/audit/timestamps.go`, and
`cmd/04-composition/main.go`:
- `Product` embeds `Timestamps` → `mug.CreatedAt` and `mug.Touch(...)` exist without
  `Product` redeclaring them.
- `PriceWatcher` holds a `notify.Sender` → it announces reprices without being a sender.

## Exercises
1. Give `PriceWatcher` an SMS sender instead of email by changing only the value passed in
   `main.go`. Confirm `PriceWatcher`'s code doesn't change.
2. Add a `Slug() string` method to `Timestamps`? No — that doesn't belong there. Instead add
   it to `Product`. Decide for two more hypothetical methods whether they belong on the
   embedded type or the outer type, and why.
3. Embed *two* types in a struct where both have a method of the same name. What happens, and
   how do you disambiguate? (This is Go's answer to multiple inheritance's "diamond problem".)

## Recap & next
You can reuse behavior without inheritance. Next: treating many concrete types uniformly.

← [03 · Abstraction](03-abstraction.md) · → [05 · Polymorphism the Go way](05-polymorphism.md)
