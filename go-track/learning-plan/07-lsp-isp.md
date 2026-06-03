# 07 · Liskov Substitution & Interface Segregation

## Goal
Understand what makes a subtype *honestly* substitutable (LSP), and how small,
client-focused interfaces (ISP) stop you from ever writing a dishonest one.

## Theory
**LSP — Liskov Substitution Principle.** "Subtypes must be substitutable for their base
types." If code works with an abstraction `T`, it must keep working when handed *any* concrete
implementation of `T` — no surprises. Violations look like:

- A method that throws "not supported" for some implementations (the canonical
  `ReadOnlyList.add()` that blows up).
- A subtype that *strengthens preconditions* (demands more than the contract) or *weakens
  postconditions* (promises less).
- The classic **Square extends Rectangle**: setting width independently of height is part of
  Rectangle's contract, but Square can't honor it — so a Square is *not* a substitutable
  Rectangle, even though "a square is a rectangle" in geometry.

LSP is the real test for "should this be inheritance/subtyping?" The question isn't "is-a"
in English; it's "does it honor every promise the supertype makes?"

**ISP — Interface Segregation Principle.** "Clients should not be forced to depend on methods
they don't use." Fat interfaces push implementers to stub out or fake methods — and that
stubbing is exactly how LSP violations are born. Split a fat interface into role-sized ones
and each implementer only promises what it can actually deliver.

The two principles are partners: **ISP prevents the situations where LSP gets violated.**

## Language notes (Go)
- Go's tiny implicit interfaces make ISP the *default*. The standard library models the
  ideal: `io.Reader`, `io.Writer`, `io.Closer` are separate one-method interfaces, composed
  (`io.ReadWriteCloser`) only when a client truly needs all three.
- Because interfaces are satisfied implicitly and declared by the consumer, a function can
  ask for *exactly* the methods it calls. A type that lacks a method simply won't satisfy the
  interface — the compiler enforces substitutability *at the boundary*.
- This turns many runtime LSP violations into **compile errors**: you can't pass a
  `*Vault` where a `Withdrawable` is required, because it has no `Withdraw`.

## Practice
```bash
docker compose run --rm golang go run ./cmd/07-lsp-isp
```

Read `internal/banking/banking.go` and `cmd/07-lsp-isp/main.go`:
- `Depositable` and `Withdrawable` are segregated (ISP) — one capability each.
- `transfer` asks for the two narrow capabilities it uses, not a fat `Account`.
- `Vault` has no `Withdraw`, so `transfer(vault, …)` won't compile. The LSP violation
  ("withdraw from a vault you can't withdraw from") is impossible to write.
- `Checking` honors the `Withdrawable` contract fully: it returns a documented error, it
  doesn't panic — so it's a faithful substitute.

## Exercises
1. Add a `FixedDeposit` account that forbids withdrawals until a maturity date. Should it
   implement `Withdrawable`? Decide using LSP, not English. (Hint: don't — give it a separate
   `MatureWithdrawable` capability.)
2. Implement the Square/Rectangle trap with structs and an interface, then show how splitting
   into `Shape{ Area() }` removes the LSP problem entirely.
3. Find a fat interface in code you know (or imagine one with 8 methods). Split it into
   role interfaces and note which clients shrink.

## Recap & next
Honest substitutes, client-sized interfaces. Next: the keystone principle the whole
architecture half of this course rests on — **Dependency Inversion**.

← [06 · SRP & OCP](06-srp-ocp.md) · → [08 · Dependency Inversion & injection](08-dip-di.md)
