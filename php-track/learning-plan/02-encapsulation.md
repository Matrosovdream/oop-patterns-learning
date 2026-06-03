# 02 · Classes, objects & encapsulation

## Goal
Build a type whose internal state can never be put into an invalid configuration from
outside — the heart of encapsulation — using `Money`.

## Theory
**Encapsulation** = bundling data with the operations allowed on it, and hiding the data so
those operations are the *only* way to change it. The payoff is the **invariant**: a rule that
is always true for a valid object ("money has exactly one currency"; "an order total is never
negative"). If state is public, every line in your program is a suspect when an invariant
breaks. If state is private and reachable only through methods, the suspect list is one class.

A well-encapsulated type exposes *intentions*, not *fields*:
- ❌ `$wallet->amount -= 475;`
- ✅ `$wallet->subtract($coffee);` — and it can refuse to mix currencies.

Encapsulation pairs naturally with **immutability**: `Money::subtract()` returns a *new* Money
instead of mutating `$this`, so a value you're holding never changes underneath you.

## Language notes (PHP)
- Visibility is per-property: `private`, `protected`, `public`. Default to `private`; widen
  only with a reason.
- **`readonly` properties** (PHP 8.1+) can be assigned once (in the constructor) and never
  again — the language enforces immutability for you. `Money`'s fields are `private readonly`.
- **Constructor property promotion** declares and assigns a property in the signature:
  `private function __construct(private readonly int $minor, ...)`.
- **Named constructors**: a `private` constructor plus `public static` factories
  (`Money::of`, `Money::fromMajor`) gives you expressive, validated construction and is the
  only door in. `final` stops a subclass from prying the door open.

## Practice
```bash
docker compose run --rm php php examples/02-encapsulation.php
```

Read `examples/02-encapsulation.php` with `src/Money/Money.php`. Trace:
1. `subtract()` returns a new value — `$wallet` is unchanged afterward.
2. `add()` with EUR throws — the type refuses to compute a meaningless result.

Now try to break it: add `$wallet->minor = 0;` to the script. You'll get a fatal error —
`minor` is private and readonly. That error *is* the encapsulation working.

## Exercises
1. Add an `allocate(int $parts): array` method that splits an amount into N parts with the
   remainder distributed so the parts sum back to the original (no lost cents). Keep fields
   private.
2. Add a `negate(): self` returning the same amount with the opposite sign.
3. Why is `Money` immutable but `Product` (lesson 04) mutable? Write one sentence.

## Recap & next
State is protected behind methods and invariants. Next: describing behavior abstractly so
callers don't depend on concrete types.

← [01 · Introduction](01-introduction.md) · → [03 · Abstraction & interfaces](03-abstraction.md)
