# Go track — learning plan

Object-oriented *design* in a language that deliberately has no classes and no
inheritance. Go gives you structs, methods, embedding, and **implicit interfaces** — and
it turns out that's enough for almost everything, often with less ceremony. Where a
pattern is "just a language feature" in Go, we'll say so.

- Code project: [`../go-app/`](../go-app) — module `oop-patterns-learning/go-app`.
- Run a lesson: `docker compose run --rm golang go run ./cmd/<NN-topic>`
  (or locally: `cd go-track/go-app && go run ./cmd/<NN-topic>`).
- Track your progress in [`PROGRESS.md`](PROGRESS.md).

## How the code is organized

```
go-app/
├── go.mod
├── cmd/<NN-topic>/main.go   # the runnable demo for each lesson
└── internal/<area>/         # the real implementations the demos import
    ├── money/               # Money value object (used everywhere)
    ├── catalog/             # products
    ├── pricing/             # discount strategies
    └── ordering/            # the Order aggregate
```

Lesson demos live in `cmd/`; reusable types live in `internal/`. That split is itself a
small architecture lesson — runnable entrypoints depend on library code, never the reverse.

## Lessons

### Part 1 · OOP foundations
- [01 · Introduction & environment](01-introduction.md)
- [02 · Structs, methods & encapsulation](02-encapsulation.md)
- [03 · Abstraction & interfaces](03-abstraction.md)
- [04 · Composition over inheritance](04-composition.md)
- [05 · Polymorphism the Go way](05-polymorphism.md)

### Part 2 · SOLID
- [06 · Single Responsibility & Open/Closed](06-srp-ocp.md)
- [07 · Liskov & Interface Segregation](07-lsp-isp.md)
- [08 · Dependency Inversion & injection](08-dip-di.md)

### Part 3 · GoF design patterns
- [09 · Creational patterns](09-creational.md)
- [10 · Structural patterns](10-structural.md)
- [11 · Behavioral patterns](11-behavioral.md)

### Part 4 · Architecture
- [12 · Layered architecture](12-layered.md)
- [13 · Ports & Adapters (Hexagonal)](13-hexagonal.md)
- [14 · Clean / Onion architecture](14-clean.md)

### Part 5 · DDD + CQRS
- [15 · Strategic DDD](15-strategic-ddd.md)
- [16 · Entities & value objects](16-entities-value-objects.md)
- [17 · Aggregates, repositories & domain events](17-aggregates-events.md)
- [18 · Application layer & use cases](18-application-layer.md)
- [19 · CQRS & event sourcing](19-cqrs.md)

### Part 6 · Capstone
- [20 · Capstone — a bounded context end to end](20-capstone.md)

## Go-specific mindset (read before lesson 01)

- **No inheritance.** There is no `extends`. You reuse behavior by *embedding* a struct
  or by holding a dependency as a field. This is composition, and it's the default.
- **Interfaces are implicit and tiny.** A type satisfies an interface just by having the
  methods — no `implements` keyword. Idiomatic interfaces have 1–3 methods and are
  declared by the *consumer*, not the producer.
- **Accept interfaces, return structs.** Functions take the narrow behavior they need and
  hand back concrete values.
- **Errors are values.** No exceptions; you return `error` and handle it explicitly.
- **The zero value should be useful** when you can manage it.

Keep these in mind and most "patterns" will feel like common sense rather than ritual.
