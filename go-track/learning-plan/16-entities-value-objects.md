# 16 · Entities & value objects

## Goal
Tell entities from value objects — the most useful distinction in tactical DDD — and know
which to reach for when modeling.

## Theory
Two kinds of domain object, distinguished by *how you decide if two of them are "the same"*:

- **Value object** — defined entirely by its attributes; has **no identity**. Two are equal
  iff their contents are equal (`$10 USD` == `$10 USD`). Value objects should be
  **immutable** (change = make a new one) and are the natural home for small domain rules
  (`Money` refuses to mix currencies; an `Email` validates its format). They make code safer
  because you can pass them around freely and they can't drift into invalid states.
- **Entity** — has a distinct **identity** that persists through change. A customer with a new
  address is the *same* customer; an order whose status went pending→paid is the *same*
  order. Equality is by ID, not by contents. Entities have a lifecycle.

Why it matters: most modeling mistakes are either (a) giving an identity to something that
doesn't need one (a fat, mutable "MoneyAmount" entity in the database) or (b) treating
something with a real lifecycle as a disposable value. Getting `Money` as a value object and
`Order` as an entity right is the difference between a model that explains the business and
one that just stores rows.

**Prefer value objects.** They carry rules, they're immutable, they're trivially testable.
Reach for an entity only when you genuinely must track a thing *through change over time*.

## Language notes (Go)
- **Value objects** are small structs with value receivers and no exported mutators — see
  `money.Money` and `domain.Line`. If every field is comparable, Go's `==` gives you value
  equality for free (`l1 == l2`); otherwise write an `Equals` method (as `Money` does, to
  control the semantics).
- **Entities** hide their identity + state behind a pointer type with methods — see
  `domain.Order`. Compare them by `ID()`, never with `==` on the pointer (that's identity of
  the *Go value*, not the *domain entity*).
- Immutability isn't enforced by the language, so you achieve it by **convention**:
  unexported fields + value receivers + "derive a new one" methods.

## Practice
```bash
docker compose run --rm golang go run ./cmd/16-entities-value-objects
```

Read `cmd/16-entities-value-objects/main.go` with `internal/money/money.go` and
`internal/ordering/domain/order.go`. Two `$10 USD` are equal (value); two orders with
identical lines but different IDs are *not* the same order (identity).

## Exercises
1. Build an `Email` value object that validates on construction and is immutable. Where do the
   validation rules belong — the constructor, or every caller?
2. Add a `Quantity` value object (a positive int that rejects zero/negative) and use it in
   `Line`. What bugs does that make impossible?
3. Is `OrderID` an entity or a value object? (It's a value object that *identifies* an
   entity.) Explain the distinction in one sentence.

## Recap & next
You can classify domain objects by identity vs value. Next: grouping them into a consistency
boundary — the aggregate — and announcing change with domain events.

← [15 · Strategic DDD](15-strategic-ddd.md) · → [17 · Aggregates, repositories & domain events](17-aggregates-events.md)
