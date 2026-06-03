# 08 · Dependency Inversion & dependency injection

## Goal
Invert the direction of source-code dependencies so high-level policy stops depending on
low-level detail — and wire the concrete pieces together at a single composition root.

## Theory
**DIP — Dependency Inversion Principle.** Two clauses:
1. High-level modules should not depend on low-level modules. Both should depend on
   *abstractions*.
2. Abstractions should not depend on details. Details should depend on abstractions.

Normally control flows downhill (policy → service → database driver) and so do source
dependencies — your business logic ends up `import`ing the Postgres library. DIP *inverts*
the source dependency: you define the abstraction (`Repository`) **in the high-level module,
in its vocabulary**, and the low-level detail (`PostgresRepository`) depends on *that*. Now
the arrow points up, toward policy. This is the principle that makes hexagonal and clean
architecture (lessons 13–14) possible.

**Dependency Injection** is the *technique* that realizes DIP: instead of a class creating
its own collaborators (`s.repo = NewPostgresRepo()` — now it depends on Postgres), it
*receives* them (usually via its constructor). DI is not a framework; it's "pass
dependencies in." A DI *container* just automates the wiring for large graphs.

**Composition root.** Somewhere, real implementations must be chosen and connected. Confine
that to one place near the program's entry point (`main`). Everywhere else depends only on
abstractions. The composition root is the *only* layer allowed to know concrete types.

## Language notes (Go)
- Go DI is almost always **plain constructor injection** — `NewService(repo, sender, clock)`.
  No annotations, no magic. The community norm is to wire by hand in `main`; for big graphs,
  tools like `google/wire` (compile-time) or Uber's `fx` (runtime) generate the wiring, but
  start by hand.
- Abstractions are interfaces *defined by the consumer* (`catalog.Repository` lives in the
  catalog package, not the database package). The database package imports `catalog`, not the
  other way around — that's the inversion, visible in the import graph.
- A function type can be an injectable dependency too: `Clock func() time.Time` lets you
  inject the wall clock in prod and a fixed time in tests, no interface needed.

## Practice
```bash
docker compose run --rm golang go run ./cmd/08-dip-di
```

Read `internal/catalog/service.go`, `internal/catalog/repository.go`, and
`cmd/08-dip-di/main.go`:
- `catalog.Service` depends on the `Repository` and `notify.Sender` *interfaces* and a
  `Clock` func — never a concrete store, mailer, or `time.Now`.
- `main` is the composition root: it picks `InMemoryRepository`, `EmailSender`, and a real
  clock, then injects them. Swap `InMemoryRepository` for a database one and `Service` does
  not change.

## Exercises
1. Write a `SliceLogger`/`RecordingSender` that captures sent messages instead of printing,
   inject it in a copy of `main`, and assert on what was "sent" — DI just made `Service`
   testable with no mocking framework.
2. Inject a *fixed* `Clock` (`func() time.Time { return someFixedTime }`) and confirm product
   timestamps become deterministic.
3. Draw the import graph of `catalog`, `notify`, and `main`. Which arrows point toward policy?
   That direction is DIP.

## Recap & next
That's all of SOLID. You now have the vocabulary the design patterns are built from. Next
part: the GoF catalog, filtered through Go.

← [07 · LSP & ISP](07-lsp-isp.md) · → [09 · Creational patterns](09-creational.md)
