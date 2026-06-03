# 05 · Polymorphism

## Goal
Treat many concrete types through one interface, and recognize that "replace a `switch`/`match`
on type with polymorphism" is the move behind half the GoF patterns.

## Theory
**Polymorphism** ("many shapes") lets one piece of code operate on values of different types
through a shared abstraction. The everyday win is deleting conditionals:

```php
// Without polymorphism — a match that grows with every new discount:
$total = match ($discount->kind) {
    'percent' => /* … */,
    'fixed'   => /* … */,
    'none'    => /* … */,
};

// With polymorphism — the object knows how to apply itself:
$total = $discount->apply($price);
```

Each new behavior becomes a **new class implementing the interface**, instead of a new branch
edited into a central `match`. That's the Open/Closed Principle (lesson 06): open to new
discounts, closed to modifying the loop. Strategy, State, Command, and Visitor (lesson 11) are
all "polymorphism aimed at a specific axis of change."

## Language notes (PHP)
- PHP supports both **subtype polymorphism** (a subclass used where the parent is expected) and
  **interface polymorphism** (any implementer used where the interface is expected). Prefer
  interface polymorphism — it doesn't drag in an inheritance hierarchy.
- The compiler-ish guarantee comes from typing to the interface: `foreach ($discounts as
  Discount $d)` and PHP guarantees `$d->apply()` exists.
- For *closed* sets of variants that are pure data (not behavior), a backed **`enum`** is often
  better than a class hierarchy — enums can even carry methods. Choose enum for "one of a fixed
  few values", polymorphic classes for "open set of behaviors."

## Practice
```bash
docker compose run --rm php php examples/05-polymorphism.php
```

Read `examples/05-polymorphism.php` and `src/Pricing/*`. One `foreach` calls `apply()`/`label()`
on four different discount classes. Adding a fifth discount means adding a class — the loop
never changes.

## Exercises
1. Add a `BundleDiscount` that only applies above a threshold price; drop it into the array.
   The loop should not change.
2. Write `bestPrice(Money $price, Discount ...$discounts): Money` returning the lowest result.
   Note it works for *any* discount.
3. Rewrite the example with a `match` on a discount "kind" string and feel how much worse it
   is. Then revert. That contrast is the whole lesson.

## Recap & next
That's the four OOP pillars in PHP: encapsulation, abstraction, composition, polymorphism.
Next part turns them into principles you apply deliberately — **SOLID**.

← [04 · Inheritance vs composition](04-inheritance-composition.md) · → [06 · SRP & OCP](06-srp-ocp.md)
