# 10 · Structural patterns

## Goal
Compose objects into larger structures with Adapter, Decorator, Facade, Composite, and
Proxy — all of which lean on Go's interfaces.

## Theory
Structural patterns are about *how objects are wired together* while keeping a stable
interface so callers don't notice the wiring.

- **Adapter** — convert one interface into another the client expects. The classic use:
  fit a third-party or legacy type into *your* interface without editing either.
- **Decorator** — wrap an object in another object of the *same* interface to add behavior
  (logging, retry, caching, a prefix). Stackable, and far more flexible than subclassing for
  every combination.
- **Facade** — a single simplified entry point over a complex subsystem. Callers get one
  method instead of orchestrating five collaborators.
- **Composite** — let a *group* of objects be treated exactly like a *single* object by
  giving the group the same interface. Natural fit for trees (menus, file systems,
  send-to-many).
- **Proxy** — a stand-in with the same interface as the real object that controls access to
  it: caching, lazy loading, access control, remoting.

**Decorator vs Proxy** (they look identical — same interface, wraps another): intent differs.
A *decorator adds behavior* the caller wants; a *proxy controls access* to the real subject
(often transparently). Same shape, different purpose.

## Language notes (Go)
- All five are trivial because **any type with the right methods satisfies the interface** —
  no adapter base class, no decorator boilerplate.
- **Decorator and Proxy** are just "a struct that holds the inner interface and implements
  the same interface" — see `PrefixSender` and `CachingRates`.
- **Composite** is often just a slice type with a method: `type SenderGroup []notify.Sender`
  with its own `Send`. The group *is* a `Sender`.
- **Adapter** is a thin struct wrapping the foreign type and exposing your method names.

## Practice
```bash
docker compose run --rm golang go run ./cmd/10-structural
```

Read `internal/structural/structural.go`:
- `SMSAdapter` makes a `LegacySMS` usable wherever a `notify.Sender` is wanted.
- `PrefixSender` (decorator) wraps a sender to tag messages; it's still a `Sender`.
- `SenderGroup` (composite) sends through many as if one.
- `CachingRates` (proxy) hits the slow backend once, then serves from cache (the run prints
  "backend called 1 time(s)" despite 3 lookups).
- `Checkout` (facade) reduces cart+discount+notify to a single `PlaceOrder`.

## Exercises
1. Write a `RetrySender` decorator that retries `Send` up to N times. Stack it with
   `PrefixSender`. Order matters — try both orderings.
2. Add a `LoggingRates` *decorator* over `RateProvider` and contrast it with the `CachingRates`
   *proxy*. Same shape — articulate the intent difference in a sentence.
3. Make `SenderGroup` keep going on error and collect all failures with `errors.Join`.

## Recap & next
You can compose objects without inheritance. Next: patterns about *behavior and
communication* between objects.

← [09 · Creational](09-creational.md) · → [11 · Behavioral patterns](11-behavioral.md)
