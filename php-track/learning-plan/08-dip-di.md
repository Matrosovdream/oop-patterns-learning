# 08 · Dependency Inversion & dependency injection

## Goal
Invert the direction of dependencies so high-level policy stops depending on low-level detail —
and wire the concrete pieces together at one composition root. (In lesson 12+ Laravel's
container automates this wiring.)

## Theory
**DIP — Dependency Inversion Principle.** Two clauses:
1. High-level modules should not depend on low-level modules. Both should depend on
   *abstractions*.
2. Abstractions should not depend on details. Details should depend on abstractions.

Normally control and source dependencies both flow downhill (policy → service → DB driver), so
your business logic ends up `use`-ing the database library. DIP *inverts* the source
dependency: you define the abstraction (`ProductRepository`) in the high-level module, in its
vocabulary, and the low-level detail (`PdoProductRepository`) depends on *that*. The arrow now
points up, toward policy. This is what makes hexagonal and clean architecture (lessons 13–14)
possible.

**Dependency Injection** is the technique: instead of a class building its own collaborators
(`$this->repo = new PdoRepo()` — now coupled to PDO), it *receives* them, usually via its
constructor. DI is not a framework; it's "pass dependencies in." A DI **container** just
automates the wiring of large graphs.

**Composition root.** Somewhere, real implementations must be chosen and connected. Confine
that to one place near the entry point. Everything else depends only on abstractions.

## Language notes (PHP)
- Constructor injection with promoted, `readonly`, interface-typed properties is the idiom:
  `public function __construct(private readonly ProductRepository $repository, ...)`.
- In plain PHP the composition root is a script (`examples/08-dip-di.php`). In **Laravel** it's
  a *service provider* that binds an interface to a concrete class in the container; the
  container then auto-injects by type-hint. Same principle, automated — you'll do exactly this
  in lesson 14+.
- Type-hint constructor params to **interfaces**, never concretes — that's what makes the class
  swappable and testable.

## Practice
```bash
docker compose run --rm php php examples/08-dip-di.php
```

Read `src/Catalog/CatalogService.php`, `ProductRepository.php`, and `examples/08-dip-di.php`:
- `CatalogService` depends on the `ProductRepository`, `Sender`, and `Clock` *interfaces* —
  never a concrete store, mailer, or `new DateTimeImmutable`.
- The script is the composition root: it picks `InMemoryProductRepository`, `EmailSender`, and
  `SystemClock`, then injects them. Swap the repository for a DB one and the service is
  untouched.

## Exercises
1. Write a `RecordingSender implements Sender` that captures messages instead of printing, inject
   it, and assert on what was "sent" — DI just made `CatalogService` testable with no mocks.
2. Inject a fixed `Clock` (returning a constant time) and confirm product timestamps become
   deterministic.
3. Sketch how Laravel's container would bind `ProductRepository` → `InMemoryProductRepository`
   so the service's constructor is auto-injected. (You'll build it for real in Part 4.)

## Recap & next
That's all of SOLID. You now have the vocabulary the design patterns are built from. Next part:
the GoF catalog in PHP.

← [07 · LSP & ISP](07-lsp-isp.md) · → [09 · Creational patterns](09-creational.md)
