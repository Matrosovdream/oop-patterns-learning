# 03 · Abstraction & interfaces

## Goal
Write code that depends on *what a thing does*, not *what it is*, using PHP interfaces — and
learn what makes an abstraction good rather than leaky.

## Theory
**Abstraction** is deciding which details a caller is allowed to care about. A good
abstraction is a promise ("you can send a message with this") that hides the mechanism (SMTP?
an HTTP API? a log line?). Program to the promise and you can change mechanisms without
touching callers.

The classic failure is the **leaky abstraction**: an interface that exposes its
implementation (`getSmtpConnection()`), re-coupling callers to the thing you tried to hide.
The fix is to name behaviors in the caller's language (`send`), not the provider's mechanism.

"Program to an interface, not an implementation" (GoF) is the one-line summary — and it's the
enabler for nearly every pattern that follows.

## Language notes (PHP)
- An `interface` declares method signatures; a class opts in with `implements Sender`. Unlike
  Go, satisfaction is **explicit and nominal** — and an interface is usually declared *near*
  its implementations, not by the consumer.
- Type-hint parameters and properties to the **interface** (`private readonly Sender $sender`),
  never the concrete class. That single habit is what lets you swap implementations.
- A class can implement **many** interfaces — useful for Interface Segregation (lesson 07).
- PHP also has **abstract classes** (interface + shared implementation). Prefer a plain
  interface unless you genuinely need to share code through inheritance; an interface keeps
  callers maximally decoupled.

## Practice
```bash
docker compose run --rm php php examples/03-abstraction.php
```

Read `src/Notify/Sender.php`, `EmailSender.php`, `SmsSender.php`, and `examples/03-abstraction.php`:
- `Broadcast::to()` takes a `Sender`. It has no idea email or SMS exists.
- Swapping email → SMS in the script is a one-line change; `Broadcast` is untouched.

## Exercises
1. Add a `SlackSender implements Sender` and broadcast through it — without editing `Broadcast`
   or any existing class.
2. Make `Broadcast::to()` accept an array of `Sender`s and send through each. (You've sketched
   the Composite pattern — lesson 10.)
3. When would you choose an `abstract class` over an `interface` here? Name a concrete reason.

## Recap & next
You can hide mechanisms behind behavior. Next: the most consequential reuse decision in PHP —
inherit, compose, or trait?

← [02 · Encapsulation](02-encapsulation.md) · → [04 · Inheritance vs composition](04-inheritance-composition.md)
