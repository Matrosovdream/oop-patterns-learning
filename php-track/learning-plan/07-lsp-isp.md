# 07 · Liskov Substitution & Interface Segregation

## Goal
Understand what makes a subtype *honestly* substitutable (LSP), and how small,
client-focused interfaces (ISP) stop you from ever writing a dishonest one.

## Theory
**LSP — Liskov Substitution Principle.** "Subtypes must be substitutable for their base types."
If code works with a type `T`, it must keep working with *any* implementation of `T` — no
surprises. Violations look like:

- A method that throws "not supported" for some implementations (the canonical
  `ReadOnlyList::add()` that blows up).
- A subtype that *strengthens preconditions* (demands more than the contract) or *weakens
  postconditions* (promises less).
- The classic **Square extends Rectangle**: a Rectangle's contract lets you set width and
  height independently; a Square can't honor that — so a Square is *not* a substitutable
  Rectangle, even though "a square is a rectangle" in geometry.

LSP is the real test for "should this be inheritance/an interface implementation?" The question
isn't "is-a" in English; it's "does it honor every promise the supertype makes?"

**ISP — Interface Segregation Principle.** "Clients should not be forced to depend on methods
they don't use." Fat interfaces push implementers to stub methods they can't honor — and that
stubbing is precisely how LSP violations are born. Split a fat interface into role-sized ones
and each implementer promises only what it can deliver. **ISP prevents the situations where LSP
gets violated.**

## Language notes (PHP)
- A class can `implements` **many** interfaces, so segregating is cheap: `Checking implements
  Depositable, Withdrawable`; `Vault implements Depositable` only.
- Typing a parameter to the *narrow* interface it needs (`transfer(Withdrawable $from,
  Depositable $to)`) turns "wrong capability" into a **`TypeError` at the boundary** instead of
  a runtime "not supported" — the LSP violation literally won't run.
- Avoid fat "god" interfaces (a `RepositoryInterface` with 15 methods). Prefer several small
  ones the framework or your code can compose.

## Practice
```bash
docker compose run --rm php php examples/07-lsp-isp.php
```

Read `src/Banking/*` and `examples/07-lsp-isp.php`:
- `Depositable` and `Withdrawable` are segregated (ISP) — one capability each.
- `transfer()` asks for the two narrow capabilities it uses, not a fat `Account`.
- `Vault` doesn't implement `Withdrawable`, so `transfer($vault, …)` is a `TypeError`. The LSP
  violation ("withdraw from a vault that can't") is impossible to write.
- `Checking::withdraw()` honors its contract: it throws a documented exception, never a vague
  failure — a faithful substitute.

## Exercises
1. Add a `FixedDeposit` that forbids withdrawals until maturity. Should it implement
   `Withdrawable`? Decide using LSP, not English. (Hint: don't — give it a `MatureWithdrawable`.)
2. Implement the Square/Rectangle trap with classes, then show how a `Shape` interface with an
   `area()` method removes the LSP problem.
3. Find (or imagine) a fat 8-method interface. Split it into roles and note which clients shrink.

## Recap & next
Honest substitutes, client-sized interfaces. Next: the keystone the whole architecture half of
this course rests on — **Dependency Inversion**.

← [06 · SRP & OCP](06-srp-ocp.md) · → [08 · Dependency Inversion & injection](08-dip-di.md)
