# 04 · Inheritance vs composition (and traits)

## Goal
Make the single biggest OOP design call deliberately: when to inherit, when to compose, and
where PHP **traits** fit — because PHP, unlike Go, gives you all three.

## Theory
Two ways to reuse behavior:

- **Inheritance** ("is-a"): `class SavingsAccount extends Account`. Powerful but rigid —
  subclasses couple to the base class's internals, deep hierarchies become fragile (the
  "fragile base class" problem), and PHP allows only one parent. Overriding can silently
  violate the base class's contract (lesson 07's Liskov test).
- **Composition** ("has-a"/"uses-a"): an object *holds* collaborators and delegates to them.
  More flexible — swap the collaborator at runtime, combine many, change one without
  disturbing others.

The industry rule, straight from the GoF book: **favor composition over inheritance.** Reserve
inheritance for genuine, stable, substitutable "is-a" relationships (rarer than they look).

PHP adds a third option: **traits** — horizontal reuse that copies methods/properties into a
class at compile time, with no "is-a" relationship and no shared base type. Great for truly
cross-cutting concerns (timestamps, soft-deletes). The trap: traits create *implicit*
coupling and can't be swapped at runtime, so they're **not** a substitute for composing a
collaborator you should be injecting. Rule of thumb: trait for a mixin of mechanical
behavior; composition for a dependency with its own identity/lifecycle.

## Language notes (PHP)
- `use SomeTrait;` inside a class pulls in its members — see `Product use Timestamps`. This is
  the closest PHP analog to Go's struct embedding.
- Multiple traits can be combined; conflicts are resolved explicitly with `insteadof`/`as`
  (PHP's answer to the multiple-inheritance "diamond problem").
- Composition is just a typed (interface) property set in the constructor: `PriceWatcher` holds
  a `Sender`. Type the property to the **interface** so it stays swappable.

## Practice
```bash
docker compose run --rm php php examples/04-inheritance-composition.php
```

Read `src/Support/Timestamps.php`, `src/Catalog/Product.php`, and `examples/04-...php`:
- `Product` *uses* the `Timestamps` trait → it has `touch()/createdAt()` without inheriting.
- `PriceWatcher` *holds* a `Sender` → it reuses sending by composition, and could take SMS
  instead of email with no change to its own code.

## Exercises
1. Swap `PriceWatcher`'s `EmailSender` for an `SmsSender` by changing only the constructor
   argument. Confirm `PriceWatcher`'s body doesn't change.
2. Suppose you were tempted to write `class PriceWatcher extends EmailSender`. List two
   concrete problems that creates, then explain why composition avoids both.
3. Two traits both define a `log()` method and you `use` both. How do you resolve the clash?

## Recap & next
You can reuse behavior three ways and know which to pick. Next: treating many concrete types
uniformly.

← [03 · Abstraction](03-abstraction.md) · → [05 · Polymorphism](05-polymorphism.md)
