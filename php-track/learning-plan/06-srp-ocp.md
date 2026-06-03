# 06 · Single Responsibility & Open/Closed

## Goal
Apply the first two SOLID principles: give each class *one reason to change* (SRP), and make
behavior extensible *without editing existing code* (OCP).

## Theory
**SRP — Single Responsibility Principle.** "A class should have one, and only one, reason to
change." Sharper: a class should answer to *one actor* — one source of change requests. A
`Cart` that totals **and** formats a receipt **and** emails it **and** saves to a DB has four
bosses (finance, design, marketing, ops); a request from any one risks breaking the others.
Split them and each change stays local.

SRP is about **cohesion**, not size. It's not "one method per class" — it's "things that
change together live together; things that change for different reasons live apart."

**OCP — Open/Closed Principle.** "Software entities should be open for extension, but closed
for modification." Add behavior by adding code, not by editing code that already works (and is
already tested). The mechanism is almost always *polymorphism behind an interface*. The
tell-tale OCP smell is a `match`/`if-else` over a type tag that grows with every new business
case — replace it with an interface + one class per case.

## Language notes (PHP)
- SRP often shows up as small, focused **classes** and free functions. Watch for the `*Manager`
  / `*Helper` / `*Service` that quietly accretes responsibilities — that's the anti-pattern.
- OCP is achieved by **typing to an interface** and accepting it as a parameter: `checkout(Cart
  $cart, Discount $discount)` is closed (you don't edit it) yet open (any `Discount` works).
- PHP's `final` keyword supports OCP from the other side: marking a class `final` says "extend
  my behavior by composing/implementing, not by subclassing me."

## Practice
```bash
docker compose run --rm php php examples/06-srp-ocp.php
```

Read `examples/06-srp-ocp.php`:
- `checkout()` (compute), `renderReceipt()` (format), and `EmailSender::send()` (deliver) are
  three responsibilities in three places — change the receipt layout without touching pricing.
- `checkout()` takes a `Discount`; adding a `BlackFridayDiscount` never edits it.

## Exercises
1. Add a new `Discount` and run it through `checkout()` *without modifying `checkout()`*. That's
   OCP; not touching tested code is the payoff.
2. `renderReceipt()` builds lines *and* formats totals. One responsibility or two? Argue both
   sides, then decide.
3. Introduce a `ReceiptStore` that persists receipts. Where does it go so `checkout()` and
   `renderReceipt()` keep their single responsibilities?

## Recap & next
One reason to change; extend without editing. Next: two principles about *substitutability*.

← [05 · Polymorphism](05-polymorphism.md) · → [07 · Liskov & Interface Segregation](07-lsp-isp.md)
