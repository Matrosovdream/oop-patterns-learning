# 02 · Structs, methods & encapsulation

## Goal
Build a type whose internal state can never be put into an invalid configuration from
outside — the heart of encapsulation — using `money.Money`.

## Theory
**Encapsulation** = bundling data with the operations allowed on it, and hiding the data so
those operations are the *only* way to change it. The payoff is the **invariant**: a rule
that is always true for a valid object (e.g. "money has exactly one currency", "an order
total is never negative"). If state is public, every line of code in your program is a
suspect when an invariant breaks. If state is private and reachable only through methods,
the suspects are just that one type.

A well-encapsulated type exposes *intentions*, not *fields*:
- ❌ `wallet.amount = wallet.amount - 475`
- ✅ `wallet.Sub(coffee)` — and it can reject mixing currencies.

Encapsulation is also where **immutability** shines: `Money.Sub` returns a *new* `Money`
instead of mutating the receiver, so a value you hold never changes under you.

## Language notes (Go)
- Visibility is per **package**, by case. In `package money`, the fields `minor` and
  `currency` are lowercase → invisible outside the package. `New`, `Add`, `Amount` are
  uppercase → exported. There is no `private`/`public` keyword and no per-field control
  finer than the package.
- A **method** is just a function with a receiver: `func (m Money) Add(...)`. A *value*
  receiver (`m Money`) gets a copy and can't mutate the original — perfect for an immutable
  value type. A *pointer* receiver (`t *Timestamps`) can mutate. Choosing value receivers
  here is what makes immutability natural.
- Go has no constructors built into the language; the convention is an exported `New...`
  function. Because the fields are unexported, that function is the only door in.

## Practice
```bash
docker compose run --rm golang go run ./cmd/02-encapsulation
```

Read `go-app/cmd/02-encapsulation/main.go` alongside `internal/money/money.go`. Trace:
1. `wallet.Sub(coffee)` returns a **new** value — `wallet` is unchanged afterward.
2. `wallet.Add(euros)` returns an **error** because the currencies differ. The type refuses
   to compute a meaningless result.

Now try to break it: in `main.go`, attempt `wallet.minor = 0`. It won't compile — `minor`
is unexported. That compiler error *is* the encapsulation working.

## Exercises
1. Add an `Allocate(parts int) []Money` method that splits an amount into N parts with the
   remainder distributed so the parts sum back to the original (no lost cents). Keep the
   fields private.
2. Add a `Negate() Money` method returning the same amount with the opposite sign.
3. Why does `Money` use a value receiver while `audit.Timestamps.Touch` (lesson 04) uses a
   pointer receiver? Write a one-sentence answer.

## Recap & next
State is now protected behind methods and invariants. Next we describe *behavior* abstractly
so callers don't depend on concrete types.

← [01 · Introduction](01-introduction.md) · → [03 · Abstraction & interfaces](03-abstraction.md)
