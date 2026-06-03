# 06 · Single Responsibility & Open/Closed

## Goal
Apply the first two SOLID principles: give each unit *one reason to change* (SRP), and
make behavior extensible *without editing existing code* (OCP).

## Theory
**SRP — Single Responsibility Principle.** "A module should have one, and only one, reason
to change." The sharper phrasing (Uncle Bob's): a module should be responsible to *one
actor* — one source of change requests. A `Cart` that totals itself **and** formats a
receipt **and** emails it **and** saves to a DB has four bosses (finance, design, marketing,
ops); a change from any of them risks breaking the others. Split those into separate units
and each change stays local.

SRP is not "one method per class" or "tiny classes everywhere" — that's shrapnel. It's about
*cohesion*: the things that change together live together; the things that change for
different reasons live apart.

**OCP — Open/Closed Principle.** "Software entities should be open for extension, but closed
for modification." You should be able to add new behavior by adding new code, not by editing
code that already works (and is already tested). The mechanism is almost always
*polymorphism behind an abstraction*: a stable interface that new implementations plug into.

The tell-tale OCP smell is a `switch`/`if-else` over a type tag that grows every time the
business adds a case. Replace it with an interface + one type per case.

## Language notes (Go)
- SRP in Go often means a small **package** or a small **function**, not a class. Go's bias
  toward many small packages and functions makes SRP feel natural.
- OCP in Go is achieved with **interfaces + functions as parameters**. `Checkout(c, d
  pricing.Discount)` is closed (you don't edit it) yet open (any new `Discount` works).
- Go has no abstract classes to extend, so OCP is *always* expressed via interfaces or
  higher-order functions — there's no inheritance temptation to get it wrong.

## Practice
```bash
docker compose run --rm golang go run ./cmd/06-srp-ocp
```

Read `cmd/06-srp-ocp/main.go`:
- `Checkout` (compute), `renderReceipt` (format), and `EmailSender.Send` (deliver) are three
  responsibilities in three places — change the receipt layout without touching pricing.
- `Checkout` takes `pricing.Discount`. Adding a `BlackFridayDiscount` never edits `Checkout`.

## Exercises
1. Add a new discount type and run it through `Checkout` *without modifying `Checkout`*. That
   it's possible is OCP; that you didn't touch tested code is the payoff.
2. The `renderReceipt` function both builds lines and formats totals. Is that one
   responsibility or two? Argue both sides, then decide.
3. Introduce a `ReceiptStore` that persists receipts. Where does it go so `Checkout` and
   `renderReceipt` keep their single responsibilities?

## Recap & next
One reason to change; extend without editing. Next: two principles about *substitutability*.

← [05 · Polymorphism](05-polymorphism.md) · → [07 · Liskov & Interface Segregation](07-lsp-isp.md)
