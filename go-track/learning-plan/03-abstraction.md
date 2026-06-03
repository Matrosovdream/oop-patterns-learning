# 03 · Abstraction & interfaces

## Goal
Write code that depends on *what a thing does*, not *what it is*, using Go's implicit
interfaces — and understand why Go interfaces are small and defined by the consumer.

## Theory
**Abstraction** is deciding which details a caller is allowed to care about. A good
abstraction is a promise ("you can send a message with this") that hides the mechanism
(SMTP? an HTTP API? a log line?). Program against the promise and you can swap mechanisms
without touching the callers.

The classic failure is the **leaky abstraction**: an interface that exposes its
implementation (`getSmtpConnection()`), so callers end up coupled to the very thing you
tried to hide. The fix is to name behaviors in the *language of the caller's need*
(`Send`), not the provider's mechanism.

## Language notes (Go)
- An interface is a set of method signatures: `type Sender interface { Send(to, msg string) error }`.
- **Satisfaction is implicit.** A type is a `Sender` simply by having a `Send` method with
  that signature. It never says `implements Sender`. This means you can write an interface
  for types you don't own (e.g. a standard-library type) after the fact.
- **Interfaces are small.** The standard library's most-used interfaces have one method
  (`io.Reader`, `io.Writer`, `fmt.Stringer`, `error`). Rob Pike's proverb: *"The bigger the
  interface, the weaker the abstraction."*
- **Define interfaces where they're used, not where types are defined.** The consumer knows
  the minimal behavior it needs. The producer just writes concrete structs. (Contrast with
  Java/PHP, where the interface usually lives next to the implementation.)
- **Accept interfaces, return structs.** Functions take the narrow interface they need and
  return concrete types callers can use fully.

## Practice
```bash
docker compose run --rm golang go run ./cmd/03-abstraction
```

Read `internal/notify/notify.go` and `cmd/03-abstraction/main.go`. Notice:
- `Broadcast` takes a `Sender`. It has no idea email or SMS exists.
- `EmailSender` and `SMSSender` never mention `Sender` — they just have a `Send` method.
- Swapping email→SMS in `main.go` is a one-line change; `Broadcast` is untouched.

## Exercises
1. Add a `SlackSender` with a `Send` method and broadcast through it — without editing
   `Broadcast` or any existing type.
2. Add a `MultiSender []Sender` whose own `Send` method forwards to each child. (You've just
   written the Composite pattern — see lesson 10.)
3. Make `Broadcast` accept the standard-library `io.Writer` as well, and write a sender that
   logs to it. What does that tell you about interface reuse?

## Recap & next
You can now hide mechanisms behind behavior. Next: how Go *reuses* code without inheritance.

← [02 · Encapsulation](02-encapsulation.md) · → [04 · Composition over inheritance](04-composition.md)
