# 11 · Behavioral patterns

## Goal
Manage how objects collaborate with Strategy, Observer, Command, Template Method, and
State — and notice how several collapse into "a function value" in Go.

## Theory
Behavioral patterns are about *responsibilities and communication* between objects.

- **Strategy** — capture interchangeable algorithms behind a common type and pick one at
  runtime. (You already used it: `pricing.Discount` is Strategy as an interface.) This is the
  cure for the growing `switch`.
- **Observer** — a subject notifies registered observers when something happens, without
  knowing who they are. The basis of event systems and the seed of *domain events* (lesson
  17).
- **Command** — turn a request into an object: enables queues, logs, retries, and **undo**.
- **Template Method** — define the skeleton of an algorithm and let specific steps vary.
  Classic OOP overrides abstract methods in subclasses; the spirit is "fixed frame, pluggable
  steps."
- **State** — when an object's behavior depends on its state, model each state as a type that
  knows which transitions are legal. Replaces a brittle `if status == ...` sprawl with an
  explicit state machine.

## Language notes (Go)
- **Strategy and Template Method are usually function values.** `ShippingCost func(int64)
  Money` is a Strategy; passing a `body func() string` into `RenderReport` is a Template
  Method. No interface, no subclass.
- **Observer** is a slice of callbacks (`[]func(Event)`); `Subscribe` appends, `Publish`
  ranges. For concurrent or buffered fan-out you'd reach for **channels** instead — Go's
  built-in answer to a lot of Observer scenarios.
- **Command** maps to a small interface (`Execute`/`Undo`) plus an invoker holding history.
- **State** maps to an interface implemented by one type per state, each returning the next
  state — the compiler then can't let you call a transition method that doesn't exist.

## Practice
```bash
docker compose run --rm golang go run ./cmd/11-behavioral
```

Read `internal/behavioral/behavioral.go`:
- **Strategy**: `flat` vs `weight` shipping chosen at runtime.
- **Observer**: two subscribers both react to one published event.
- **Command**: two deposits, then `UndoLast()` rolls one back (150 → 100).
- **Template**: `RenderReport` frames a varying body.
- **State**: shipping before paying is rejected; pay → ship works; paying a shipped order is
  rejected. Illegal transitions are errors, not silent no-ops.

## Exercises
1. Add a `Cancel()` transition to the order state machine: legal from `pending`/`paid`,
   illegal from `shipped`. Notice you change states, not a central `switch`.
2. Re-implement `Observer` with a channel and a goroutine consumer. When is the channel
   version better than the callback version?
3. Add a `Withdraw` command with its own `Undo` and verify undo of a mixed sequence.

## Recap & next
That's the GoF tour. You've seen that patterns are *named solutions to recurring design
tensions* — and that the right language can make some of them disappear. Next part zooms out
from objects to whole-application **architecture**.

← [10 · Structural](10-structural.md) · → [12 · Layered architecture](12-layered.md)
