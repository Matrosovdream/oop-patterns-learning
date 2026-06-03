# 11 · Behavioral patterns

## Goal
Manage how objects collaborate with Strategy, Observer, Command, Template Method, and State —
and notice which collapse into "a callable" in modern PHP.

## Theory
Behavioral patterns are about *responsibilities and communication* between objects.

- **Strategy** — interchangeable algorithms behind a common type, chosen at runtime. (You
  already used it: `Discount` is Strategy as an interface.) The cure for the growing `match`.
- **Observer** — a subject notifies registered observers when something happens, without
  knowing who they are. The basis of event systems and the seed of *domain events* (lesson 17).
- **Command** — turn a request into an object: enables queues, logs, retries, and **undo**.
- **Template Method** — define the skeleton of an algorithm and let specific steps vary.
- **State** — when behavior depends on state, model each state as a type that knows its legal
  transitions. Replaces brittle `if ($status === …)` chains with an explicit state machine.

## Language notes (PHP)
- **Strategy and Template Method** can be plain **callables/closures** when the behavior is a
  single function (`fn (int $g) => …` for shipping; a `callable $body` for the report). Use a
  full interface when the strategy has multiple methods or needs state.
- **Observer** is typically `callable` subscribers stored in an array. Laravel formalizes this
  as events + listeners; the principle is identical.
- **Command** maps to a small `execute()`/`undo()` interface plus an invoker holding history.
  Laravel's queued jobs are commands.
- **State** maps to one class per state implementing a shared interface, each transition
  returning the next state or throwing — so an impossible transition is a thrown exception, not
  a silently-wrong status.
- For *closed value sets* (not behaviors), reach for a backed **`enum`** instead of state
  classes.

## Practice
```bash
docker compose run --rm php php examples/11-behavioral.php
```

Read `examples/11-behavioral.php` and `src/Behavioral/*`:
- **Strategy** (closures): `flat` vs `weight` shipping chosen at runtime.
- **Observer**: two subscribers react to one event.
- **Command**: two deposits, then `undoLast()` rolls one back (150 → 100).
- **Template**: `renderReport` frames a varying body.
- **State**: shipping before paying is rejected; pay → ship works; paying a shipped order is
  rejected.

## Exercises
1. Add a `cancel()` transition: legal from pending/paid, illegal from shipped. You add a method
   to each state class, not a branch to a central `switch`.
2. Re-implement Observer using Laravel-style event/listener naming (a class per event, a class
   per listener). What did formalizing it buy you?
3. Add a `Withdraw` command with its own `undo()` and verify undo of a mixed sequence.

## Recap & next
That's the GoF tour. Patterns are *named solutions to recurring design tensions* — and modern
PHP makes some of them a one-line closure. Next part zooms out to whole-application
**architecture**, where we move to Laravel.

← [10 · Structural](10-structural.md) · → [12 · Layered architecture](12-layered.md)
