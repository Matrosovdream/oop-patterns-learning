# 15 · Strategic DDD

## Goal
Learn to carve a system into **bounded contexts** with their own **ubiquitous language**, and
to map the relationships between them — the strategic decisions you make *before* writing
tactical code.

## Theory
DDD has two halves. **Strategic** DDD (this lesson) is the big picture: where the boundaries go
and what words mean inside them. **Tactical** DDD (lessons 16–19) is the building blocks inside
a boundary.

- **Ubiquitous language** — one shared vocabulary used identically by domain experts, code, and
  conversation. If the business says "place an order", the method is `place()`, not `submit()`
  or `createOrderRecord()`. Mismatched language is where bugs and misunderstandings breed.
- **Bounded context** — an explicit boundary within which a model and its language are
  consistent. The same word means *different things* across contexts: a "Product" in Catalog
  (name, images) is not the "Product" in Pricing (cost basis, tax class) nor the "Line" in
  Ordering (sku, quantity, price-at-time-of-order). Forcing one `Product` class to serve all
  three is the classic road to an unmaintainable god-model.
- **Context map** — how contexts relate: *Shared Kernel*, *Customer/Supplier*, *Conformist*,
  *Anti-Corruption Layer* (ACL), *Open Host Service*. It documents who depends on whom and how.

The deepest strategic move is **finding the boundaries** — along business capabilities, not
database tables. Get this wrong and no amount of clean tactical code saves you.

## Language notes (PHP / Laravel)
- A bounded context maps cleanly to a **namespace/folder** under `app/` (`App\Ordering\…`). A
  larger system might promote each to its own Laravel module/package or even service.
- **Shared Kernel**: `App\Shared\Money` is a value object shared by agreement across contexts —
  small, stable, and jointly owned. Keep shared kernels tiny; they couple the contexts that
  share them.
- A common Laravel anti-pattern is one giant `App\Models` namespace where every Eloquent model
  knows every other — the opposite of bounded contexts. Splitting by context fixes it.

## Practice
Sketch ShopKit's contexts and map (no code to run here — strategy is a whiteboard activity):

- **Catalog** — language: Product, SKU, Category. Owns "what we sell".
- **Pricing** — language: Money, Discount, Tax. Owns "what it costs".
- **Ordering** — language: Order, Line, OrderId, placed/paid/shipped. Owns "the purchase
  lifecycle". (This is the context you've been building under `app/Ordering`.)

Context map: `Ordering --uses--> Pricing` (Customer/Supplier); `Ordering --ACL--> Catalog`
(translates products in, lesson 18). `Money` is a Shared Kernel.

## Exercises
1. Write the ubiquitous-language glossary for Ordering: every term (Order, Line, place, pay,
   ship…) with a one-line definition. Would a domain expert agree with each?
2. ShopKit adds "subscriptions". New bounded context, or part of Ordering? Argue with language
   and rate-of-change.
3. Draw the context map and label each relationship.

## Recap & next
You can split a system along meaning, not tables. Next: the tactical building blocks inside a
context, starting with entity vs value object.

← [14 · Clean architecture](14-clean.md) · → [16 · Entities & value objects](16-entities-value-objects.md)
