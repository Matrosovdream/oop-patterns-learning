# 12 · Layered architecture

## Goal
Organize a whole application into layers with a single rule about which way dependencies
point — the foundation every later architecture refines.

## Theory
Zoom out from classes to the application. The oldest organizing idea is **layers**, each
with one job:

- **Presentation** — talks to the outside (HTTP handler, CLI, a `main`). Translates requests
  into calls inward.
- **Application** — *use cases*: orchestrates a unit of work (load, call domain, save). No
  business rules of its own.
- **Domain** — the business model and its rules. The reason the software exists.
- **Infrastructure** — technical detail: databases, message buses, email, the framework.

The rule that makes layering worth anything: **dependencies point in one direction — toward
the domain.** Presentation depends on application, application on domain; infrastructure
plugs in at the edge. The domain depends on *nothing* outward. Break that rule (let the
domain `import` your ORM) and the layers are just folders.

A common mistake is the **anemic** split where "domain" objects are bags of getters/setters
and all logic sits in the application/"service" layer. That's procedural code in costume —
the domain layer should hold behavior (lessons 16–17).

## Language notes (Go)
- Layers map naturally to **packages**: `internal/ordering/domain`, `.../app`, `.../infra`,
  and a `cmd/` for presentation. Go's import graph makes the dependency direction *visible
  and enforceable* — if `domain` imported `infra` you'd often get an import cycle, and the
  build would fail.
- `internal/` adds a hard wall: nothing outside the module can import these packages at all.
- There's no framework dictating the layout — you choose it, and the compiler keeps you honest.

## Practice
```bash
docker compose run --rm golang go run ./cmd/12-layered
```

Read, in dependency order, `internal/ordering/domain/order.go` →
`internal/ordering/app/commands.go` → `internal/ordering/infra/infra.go` →
`cmd/12-layered/main.go`. Confirm with your eyes: `domain/order.go` imports only `money`. The
`cmd` is the only place that names concrete infrastructure.

## Exercises
1. Try to make `domain` import `infra` (e.g. log from inside `Order.Place`). Notice the design
   pressure — and write down why the domain shouldn't know about logging.
2. Add a second presentation entry point (another `cmd/`) that reuses the *same* application
   handler. How much code did you duplicate? (Should be ~none.)
3. List which layer each of these belongs in: "an order can't ship unpaid", "send a
   confirmation email", "parse JSON from the request", "save the order".

## Recap & next
Layers + a one-way dependency rule. Next we sharpen "infrastructure plugs in at the edge"
into an explicit pattern: ports & adapters.

← [11 · Behavioral](11-behavioral.md) · → [13 · Ports & Adapters (Hexagonal)](13-hexagonal.md)
