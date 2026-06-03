# 01 · Introduction & environment

## Goal
Get the Go practice environment running and understand the mindset shift this track
asks of you: **design without classes or inheritance.**

## Theory
This course teaches object-oriented *design* — SOLID, the GoF patterns, layered/hexagonal
architecture, DDD. Most of that literature assumes a classical OOP language (Java, C#, PHP)
with classes, `extends`, and `implements`. Go has none of those keywords.

That's not a limitation to work around; it's a feature. Go strips OOP down to the two
things that actually carry their weight:

- **Composition** — building behavior by combining small pieces (embedding a struct, or
  holding another type as a field).
- **Interfaces** — describing *what a thing can do* with a tiny set of methods, satisfied
  implicitly.

Everything else (inheritance hierarchies, abstract base classes, `protected`, overriding)
turns out to be optional. Going through the same curriculum in Go is the fastest way to
feel which "OOP rules" are essential and which are just one language's ceremony.

## Language notes (Go)
- Code is organized into **packages** (a folder of `.go` files), not classes. A package is
  the unit of encapsulation: identifiers starting with a Capital letter are exported
  (public); lowercase ones are package-private.
- `cmd/<name>/main.go` files are runnable programs. `internal/<area>/` holds library code
  that the commands import. `internal/` is special: it can only be imported from within
  this module — a compiler-enforced architecture boundary.
- The module path is `oop-patterns-learning/go-app` (see `go.mod`), so imports look like
  `oop-patterns-learning/go-app/internal/money`.

## Practice
From the repo root:

```bash
docker compose build                                  # one-time, builds all images
docker compose run --rm golang go run ./cmd/01-intro
```

(Or locally, if you have Go 1.26: `cd go-track/go-app && go run ./cmd/01-intro`.)

Open and read these files:
- `go-app/cmd/01-intro/main.go` — the program you just ran.
- `go-app/internal/money/money.go` — the `Money` type it uses. Don't worry about the
  details yet; just notice the fields are lowercase (hidden) and you build one with
  `money.FromMajor(...)`.

You should see a list price and a tax-included price printed.

## Exercises
1. Change the tax line in `cmd/01-intro/main.go` to compute 20% VAT instead of 8% and
   re-run it.
2. Add a second product price and print it. Notice you never touch `money.go` to do this.
3. Run `go doc ./internal/money` and read the generated documentation for the package.

## Recap & next
You have a working Go environment and the right mental model: **composition + interfaces,
not inheritance.** Next we make state safe.

→ [02 · Structs, methods & encapsulation](02-encapsulation.md)
