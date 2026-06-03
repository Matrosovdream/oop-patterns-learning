# Master learning plan

This is the index for the whole course. The actual lessons live in each track:

- **PHP / Laravel** → [`../php-track/learning-plan/`](../php-track/learning-plan/README.md)
- **Go** → [`../go-track/learning-plan/`](../go-track/learning-plan/README.md)

The lesson numbers are aligned across tracks, so you can study a topic in one language
or both. Below is the shared spine and *why* each part exists.

---

## Who this is for

You can already write code and use objects/structs, but you want to design software that
stays changeable as it grows. You've heard "SOLID", "hexagonal", "aggregate", "CQRS" and
want them to stop being buzzwords. By the end you'll be able to take a fuzzy business
requirement and turn it into a clean, layered, testable design — in PHP or Go.

## How to use it

1. Build the environments once: `docker compose build` (from the repo root).
2. Pick a track. Open its `learning-plan/README.md` and start at lesson `01`.
3. For each lesson: read the theory, then run the practice command, then do the exercises.
4. Tick lessons off in that track's `PROGRESS.md`.
5. Stuck or curious? Ask the tutor "why is it done this way?" — the *why* is the point.

---

## The spine

### Part 1 · OOP foundations (01–05)
The four pillars done right, and the single most important design decision in OOP:
**composition over inheritance**. Go forces this lesson on you; PHP lets you choose
badly, so we make the choice explicit.

- `01` Introduction & environment
- `02` Classes, objects & encapsulation
- `03` Abstraction & interfaces
- `04` Inheritance vs composition
- `05` Polymorphism

### Part 2 · SOLID (06–08)
Five principles that are really one idea: *depend on abstractions, isolate the reasons a
class changes.* The payoff principle is **DIP** + dependency injection, which the entire
architecture half of the course is built on.

- `06` Single Responsibility & Open/Closed
- `07` Liskov Substitution & Interface Segregation
- `08` Dependency Inversion & dependency injection

### Part 3 · GoF design patterns (09–11)
The classic catalog — but filtered through *when you'd actually use each one* and how it
looks in a language without inheritance. Several "patterns" turn out to be language
features in Go (Strategy = a function, Singleton = `sync.Once`).

- `09` Creational — Factory Method, Abstract Factory, Builder, Singleton
- `10` Structural — Adapter, Decorator, Facade, Composite, Proxy
- `11` Behavioral — Strategy, Observer, Command, Template Method, State

### Part 4 · Architecture (12–14)
Zoom out from classes to systems. The same dependency-inversion idea, applied at the
scale of an application, gives you layered → hexagonal → clean architecture.

- `12` Layered architecture & separation of concerns
- `13` Ports & Adapters (Hexagonal)
- `14` Clean / Onion architecture & the dependency rule

### Part 5 · Domain-Driven Design + CQRS (15–19)
Where design meets the business. Strategic DDD (how to *carve up* a system) and tactical
DDD (entities, value objects, aggregates, repositories, domain events), then CQRS and a
taste of event sourcing.

- `15` Strategic DDD — ubiquitous language, bounded contexts, context mapping
- `16` Tactical DDD — entities & value objects
- `17` Aggregates, repositories, domain services & domain events
- `18` Application layer, use cases & the anti-corruption layer
- `19` CQRS & an introduction to event sourcing

### Part 6 · Capstone (20)
Put it together: one bounded context of **ShopKit**, modeled and built end to end with
everything you've learned.

- `20` Capstone

---

## The running example: ShopKit

From the patterns lessons onward, examples share a small e-commerce domain so you see how
isolated patterns combine:

- **Catalog** — products, categories.
- **Pricing** — money, discounts, tax (great for value objects + Strategy).
- **Ordering** — carts, orders, order lines (the main aggregate).

Keeping one domain means lesson 17's `Order` aggregate is the same `Order` you met as a
plain class in lesson 02 — refined, not reinvented.
