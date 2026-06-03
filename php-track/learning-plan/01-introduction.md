# 01 · Introduction & environment

## Goal
Get the PHP practice environment running and set expectations: this track teaches OOP
*design*, using modern PHP 8.3 deliberately and idiomatically.

## Theory
This course is about object-oriented *design* — SOLID, the GoF patterns, layered/hexagonal
architecture, DDD. PHP is a full classical-OOP language: classes, interfaces, abstract
classes, traits, inheritance, enums, `readonly`. That power is a double-edged sword — it's
easy to reach for inheritance or a giant "Manager" class when a smaller, composed design
would age far better. So alongside *how* to use each tool, we'll be explicit about *when not
to*.

Lessons 01–11 use a **plain-PHP** project (`php-app`) so you learn the patterns without
framework noise. Lessons 12–20 move to a **Laravel** project (`laravel-app`) where layering
and DDD earn their keep.

## Language notes (PHP)
- Every file starts with `declare(strict_types=1);` — PHP then enforces type hints instead of
  silently coercing. Treat it as non-negotiable.
- Classes are autoloaded via **Composer PSR-4**: `App\Money\Money` ⇒ `src/Money/Money.php`.
  One class/interface/trait per file, namespace matching the folder.
- We use modern features throughout: constructor property promotion, `readonly`, `enum`,
  `match`, named arguments. If any are unfamiliar, look them up as they appear — they're the
  current idiom, not exotica.

## Practice
From the repo root:

```bash
docker compose build                                   # one-time
docker compose run --rm php php examples/01-intro.php
```

(Or locally with PHP 8.3: `cd php-track/php-app && composer install -q && php examples/01-intro.php`.)

Open and read:
- `php-app/examples/01-intro.php` — the script you ran.
- `php-app/src/Money/Money.php` — the value object it uses. Note the `private` constructor and
  `readonly` properties; you build one with `Money::fromMajor(...)`.

## Exercises
1. Change the tax line in `01-intro.php` to compute 20% VAT and re-run.
2. Add and print a second product price — without editing `Money.php`.
3. Skim `Money.php`. Why is the constructor `private` and the class `final`? (We'll answer it
   properly in lesson 02 — form a guess now.)

## Recap & next
Environment works; mindset set: use PHP's power, but prefer the smaller design. Next: making
state safe.

→ [02 · Classes, objects & encapsulation](02-encapsulation.md)
