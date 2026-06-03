# 12 · Layered architecture

> From here on we work in the **Laravel** app (`php-track/laravel-app`), where layering and
> DDD start to pay for themselves.

## Goal
Organize a whole application into layers with a single rule about which way dependencies
point — the foundation every later architecture refines — inside a real Laravel project.

## Theory
Zoom out from classes to the application. The oldest organizing idea is **layers**, each with
one job:

- **Presentation** — talks to the outside (an HTTP controller, an Artisan command). Translates
  requests into calls inward.
- **Application** — *use cases*: orchestrates a unit of work (load, call domain, save). No
  business rules of its own.
- **Domain** — the business model and its rules. The reason the software exists.
- **Infrastructure** — technical detail: the database, mail, queues, the framework.

The rule that makes layering worth anything: **dependencies point one way — toward the
domain.** Presentation → application → domain; infrastructure plugs in at the edge. The domain
depends on *nothing* outward (not even Laravel). Break that — let an entity extend Eloquent's
`Model` and reach the database — and the layers are just folders.

A common mistake is the **anemic domain model**: "domain" classes are bags of getters/setters
and all logic sits in fat "services". That's procedural code in costume. Our `Order` holds
behavior (lessons 16–17); the application layer only choreographs.

## Language notes (PHP / Laravel)
- Layers map to namespaces under `app/`: `App\Ordering\Domain`, `App\Ordering\Application`,
  `App\Ordering\Infrastructure`, with controllers/commands as presentation.
- The default Laravel instinct is "fat controllers + Eloquent everywhere". This track
  deliberately keeps the **domain free of Eloquent**: the `Order` aggregate is a plain PHP
  object; persistence is an Infrastructure concern behind a repository (lesson 17).
- Laravel's **service container** wires the layers (lesson 14) — but the dependency *direction*
  is your discipline, not the framework's: nothing in `Domain` may `use` anything from
  `Infrastructure` or `Illuminate\*`.

## Practice
```bash
docker compose run --rm laravel php artisan ordering:demo
# or: docker compose up laravel  → open http://localhost:8000/ordering
```

Read, in dependency order, `app/Ordering/Domain/Order.php` → `app/Ordering/Application/
PlaceOrderHandler.php` → `app/Ordering/Infrastructure/InMemoryOrderRepository.php` →
`app/Http/Controllers/OrderingController.php`. Confirm with your eyes: `Order.php` imports only
`App\Shared\Money` and its own domain — no `Illuminate`, no database.

## Exercises
1. Grep `app/Ordering/Domain` for `Illuminate` and `Eloquent`. Zero hits is the win — explain
   why that matters when Laravel releases a new major version.
2. The same use cases back both `OrderingController` (HTTP) and `OrderingDemo` (CLI). How much
   logic is duplicated between them? (Should be ~none.)
3. Classify each into a layer: "an order can't ship unpaid", "send a confirmation email",
   "validate the request JSON", "save the order".

## Recap & next
Layers + a one-way dependency rule, in Laravel. Next we sharpen "infrastructure plugs in at the
edge" into an explicit pattern: ports & adapters.

← [11 · Behavioral](11-behavioral.md) · → [13 · Ports & Adapters (Hexagonal)](13-hexagonal.md)
