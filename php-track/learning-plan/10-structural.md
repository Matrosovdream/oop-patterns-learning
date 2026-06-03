# 10 · Structural patterns

## Goal
Compose objects into larger structures with Adapter, Decorator, Facade, Composite, and Proxy
— all behind a stable interface.

## Theory
Structural patterns are about *how objects are wired together* while keeping a stable
interface so callers don't notice the wiring.

- **Adapter** — convert one interface into another the client expects. The classic use: fit a
  third-party or legacy class into *your* interface without editing either.
- **Decorator** — wrap an object in another of the *same* interface to add behavior (logging,
  retry, caching, a tag). Stackable, and far more flexible than a subclass per combination.
- **Facade** — a single simplified entry point over a complex subsystem.
- **Composite** — let a *group* of objects be treated like a *single* object by giving the
  group the same interface. Natural for trees (menus, send-to-many).
- **Proxy** — a stand-in with the same interface that controls access to the real object:
  caching, lazy loading, access control, remoting.

**Decorator vs Proxy** look identical (same interface, wraps another). Intent differs: a
*decorator adds behavior* the caller wants; a *proxy controls access* to the real subject.

## Language notes (PHP)
- Adapter, Decorator, Composite, and Proxy here all `implements App\Notify\Sender` (or
  `RateProvider`) and hold the wrapped instance as a constructor-injected property. Type the
  property to the interface so wrappers nest freely.
- **Composite** takes the children as a variadic interface param (`Sender ...$children`).
- The `??=` operator makes the **Proxy** cache a one-liner (`$this->cache[$k] ??= $this->inner->
  rate($k)`).
- Laravel itself is full of these: facades (the literal `Facade` class), the container as a
  kind of proxy/factory, middleware as decorators around the request handler.

## Practice
```bash
docker compose run --rm php php examples/10-structural.php
```

Read `src/Structural/*`:
- `SmsAdapter` makes `LegacySms` usable wherever a `Sender` is expected.
- `PrefixSender` (decorator) tags messages; it's still a `Sender`.
- `SenderGroup` (composite) sends through many as if one.
- `CachingRates` (proxy) hits the slow backend once (the run prints "called 1 time(s)").
- `Checkout` (facade) reduces cart+discount+notify to one `placeOrder()`.

## Exercises
1. Write a `RetrySender` decorator that retries `send()` up to N times. Stack it with
   `PrefixSender`; try both orderings.
2. Add a `LoggingRates` *decorator* over `RateProvider` and contrast it with the `CachingRates`
   *proxy*. Same shape — state the intent difference in one sentence.
3. Find three structural patterns in Laravel's own source or docs (hint: facades, middleware,
   the container).

## Recap & next
Compose objects without inheritance. Next: patterns about *behavior and communication*.

← [09 · Creational](09-creational.md) · → [11 · Behavioral patterns](11-behavioral.md)
