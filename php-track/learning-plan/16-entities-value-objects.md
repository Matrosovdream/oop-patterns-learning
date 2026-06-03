# 16 · Entities & value objects

## Goal
Tell entities from value objects — the most useful distinction in tactical DDD — and model
each correctly in PHP.

## Theory
Two kinds of domain object, distinguished by *how you decide if two are "the same"*:

- **Value object** — defined entirely by its attributes; **no identity**. Two are equal iff
  their contents are equal (`$10 USD` equals `$10 USD`). Value objects should be **immutable**
  (change = make a new one) and are the natural home for small domain rules (`Money` refuses to
  mix currencies; an `Email` validates its own format). They make code safer: pass them around
  freely, they can't drift into an invalid state.
- **Entity** — has a distinct **identity** that persists through change. A customer with a new
  address is the *same* customer; an order that went pending→paid is the *same* order. Equality
  is by id, not contents. Entities have a lifecycle.

Most modeling mistakes are (a) giving an identity to something that needs none (a mutable
"MoneyAmount" row) or (b) treating something with a real lifecycle as a throwaway value.
Getting `Money` as a value object and `Order` as an entity is the difference between a model
that explains the business and one that just stores rows.

**Prefer value objects.** They carry rules, they're immutable, they're trivial to test. Reach
for an entity only when you must track a thing *through change over time*.

## Language notes (PHP)
- **Value objects**: `final` classes with `private`/`public readonly` fields, named
  constructors, and an `equals()` method for value comparison — see `App\Shared\Money`,
  `Ordering\Domain\OrderId`, and `Ordering\Domain\Line`.
- **Entities**: a class with a private identity + state and methods — see `Ordering\Domain\
  Order`. Compare by `id()->equals(...)`, never by `===` on the object (that's PHP-object
  identity, not domain identity).
- **`readonly`** gives you compiler-enforced immutability for value objects (PHP 8.1+). Use a
  backed **`enum`** for closed value sets (`OrderStatus`) instead of string constants.
- Beware Eloquent here: an `Eloquent\Model` is a mutable, identity-bearing, DB-coupled object —
  fine as a *persistence* record, wrong as a *value object*. Keep VOs as plain PHP.

## Practice
```bash
docker compose run --rm laravel php artisan ordering:demo
```

Read `app/Shared/Money.php`, `app/Ordering/Domain/OrderId.php`, `Line.php`, and `Order.php`:
- `Money`/`Line`/`OrderId` are value objects (immutable, equality by value).
- `Order` is an entity: identity in `OrderId`, mutable status, equality by id.

## Exercises
1. Build an `Email` value object that validates in its constructor and is immutable. Where do
   the validation rules belong — the constructor, or every caller?
2. Add a `Quantity` value object (a positive int rejecting zero/negative) and use it in `Line`.
   What bugs does that make impossible?
3. `OrderId` — entity or value object? (A value object that *identifies* an entity.) Explain in
   one sentence.

## Recap & next
You can classify domain objects by identity vs value. Next: grouping them into a consistency
boundary — the aggregate — and announcing change with domain events.

← [15 · Strategic DDD](15-strategic-ddd.md) · → [17 · Aggregates, repositories & domain events](17-aggregates-events.md)
