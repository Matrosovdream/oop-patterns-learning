# PHP / Laravel track — learning plan

Object-oriented design in a classical-OOP language. PHP gives you the full toolkit —
classes, interfaces, abstract classes, traits, inheritance, enums, readonly properties — so
this track is where you learn not just *how* to use those tools but *when not to* (when does
inheritance bite? when is a trait the wrong reuse mechanism?).

This track uses **two** code projects:

| Lessons | Project | Run with | Why |
|---------|---------|----------|-----|
| 01–11 (OOP, SOLID, GoF) | [`../php-app/`](../php-app) | `docker compose run --rm php php examples/<NN>.php` | pure PHP — patterns without framework noise |
| 12–20 (architecture, DDD, CQRS) | [`../laravel-app/`](../laravel-app) | `docker compose up laravel` then hit a route / run an artisan command | a real framework, where layering and DDD pay off |

Track your progress in [`PROGRESS.md`](PROGRESS.md).

## How the plain-PHP code is organized (`php-app`)

```
php-app/
├── composer.json          # PSR-4: App\ => src/
├── src/<Area>/*.php        # reusable classes (Money, Notify, Pricing, …)
└── examples/<NN>-*.php     # one runnable script per lesson
```

Run an example (the entrypoint installs the autoloader on first run):

```bash
docker compose run --rm php php examples/01-intro.php
# or locally: cd php-track/php-app && composer install -q && php examples/01-intro.php
```

## How the Laravel code is organized (`laravel-app`)

A standard Laravel app, plus a DDD-layered bounded context under `app/`:

```
laravel-app/app/
├── Ordering/
│   ├── Domain/            # Order aggregate, value objects, events, repository interface (port)
│   ├── Application/       # use cases (command/query handlers), DTOs
│   └── Infrastructure/    # Eloquent/in-memory repositories, event publishers (adapters)
└── Providers/             # binds ports → adapters in the service container
```

## Lessons

### Part 1 · OOP foundations
- [01 · Introduction & environment](01-introduction.md)
- [02 · Classes, objects & encapsulation](02-encapsulation.md)
- [03 · Abstraction & interfaces](03-abstraction.md)
- [04 · Inheritance vs composition (and traits)](04-inheritance-composition.md)
- [05 · Polymorphism](05-polymorphism.md)

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

## PHP-specific mindset (read before lesson 01)

- **You have inheritance — use it sparingly.** PHP lets you `extends`, but deep hierarchies
  age badly. This track keeps inheritance for genuine "is-a substitutable" cases and reaches
  for composition / interfaces / traits otherwise (lesson 04 makes the call explicit).
- **Lean on modern PHP.** `readonly` properties for immutability, enums for closed sets of
  values, constructor property promotion, `match`, named arguments, first-class callables.
- **Interfaces are explicit** (`implements`) and usually declared near their implementations —
  the opposite of Go's consumer-defined interfaces. Both work; the difference is instructive.
- **Type everything.** `declare(strict_types=1)` at the top of every file; type-hint
  parameters and returns. The compiler-ish guarantees you get are most of the safety.
