# 05 · Polymorphism the Go way

## Goal
Treat many different concrete types through one interface, and recognize that "replace a
`switch` on type with polymorphism" is the move behind half the GoF patterns.

## Theory
**Polymorphism** ("many shapes") lets one piece of code operate on values of different
types through a shared abstraction. The everyday win is deleting conditionals:

```go
// Without polymorphism — a switch that grows every time a discount is added:
switch d.kind {
case "percent": ...
case "fixed":   ...
case "none":    ...
}

// With polymorphism — the type knows how to apply itself:
price = d.Apply(price)
```

Each new behavior becomes a new type implementing the interface, instead of a new branch
edited into a central `switch`. That is literally the Open/Closed Principle (lesson 06): open
to new discounts, closed to modification of the loop. Strategy, State, Command, and Visitor
(lessons 11) are all "polymorphism applied to a specific axis of change".

## Language notes (Go)
- Go has **no inheritance-based (subtype) polymorphism** — only **interface polymorphism**.
  A `[]pricing.Discount` can hold a `NoDiscount`, a `PercentageOff`, etc., because each
  satisfies the interface. At the call site Go dispatches to the right concrete method
  dynamically (via an interface's internal type+method pointer).
- There are **no generics needed** for this kind of polymorphism — interfaces handle "different
  types, same behavior". Generics (`[T any]`) solve a different problem: "same behavior, but
  keep the concrete type", e.g. a typed container. Don't reach for generics where an
  interface is clearer.
- A nil interface vs an interface holding a nil pointer is a classic Go gotcha — be careful
  returning concrete nil pointers as interfaces.

## Practice
```bash
docker compose run --rm golang go run ./cmd/05-polymorphism
```

Read `internal/pricing/discount.go` and `cmd/05-polymorphism/main.go`. One `for` loop calls
`Apply`/`Label` on four different concrete discount types. Adding a fifth discount means
adding a type — the loop never changes.

## Exercises
1. Add a `BundleDiscount` that applies percentage off only above a threshold price; drop it
   into the slice in `main.go`. The loop should not change.
2. Write a function `BestPrice(price money.Money, ds ...pricing.Discount) money.Money` that
   returns the lowest price across several discounts. Note it works for *any* discount.
3. Rewrite `main.go` to instead use a `switch` on a discount "kind" string and feel how much
   worse it is. Then revert. This contrast is the whole lesson.

## Recap & next
That's the four OOP pillars in Go: encapsulation, abstraction, composition, polymorphism.
Next part turns them into principles you can apply deliberately — **SOLID**.

← [04 · Composition](04-composition.md) · → [06 · SRP & OCP](06-srp-ocp.md)
