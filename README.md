# OOP, Design Patterns, Architecture & DDD — a hands-on course

A personal, self-paced course for learning **object-oriented design** the way it's
actually used in production: from OOP fundamentals and **SOLID**, through the **GoF
design patterns**, into **software architecture** (layered, hexagonal, clean), and
finally **Domain-Driven Design + CQRS**.

Every topic is taught **twice** — once for **PHP / Laravel** and once for **Go** —
because the same idea looks different in a classical-OOP language and in a
composition-first one. Seeing both is the fastest way to understand what the pattern
*really* is versus what's just language ceremony.

> Modeled on the structure of my `vuejs-learning` repo: a `learning-plan/` of numbered
> lessons (theory) paired with runnable practice projects. Claude is the tutor — ask
> questions as you go.

---

## The two tracks

| Track | Folder | Practice runs in | Teaches the topic using… |
|-------|--------|------------------|--------------------------|
| **PHP / Laravel** | [`php-track/`](php-track/) | `php` + `laravel` containers | classes, interfaces, traits, a DI container, Eloquent |
| **Go** | [`go-track/`](go-track/) | `golang` container | structs, embedding, implicit interfaces, packages |

Each track has its own `learning-plan/` (a `README.md` index, a `PROGRESS.md` tracker,
and 20 numbered lesson files) plus a code project the lessons drive.

Pick one track and go top-to-bottom, or learn a topic in both and compare — the
lesson numbers line up (lesson `09` is GoF creational patterns in *both* tracks).

Start here: **[learning-plan/README.md](learning-plan/README.md)** — the master index.

---

## The three practice environments (Docker)

You don't need PHP or Go installed locally — only Docker. Three Compose services map
to the three code projects:

| Service | Image | Code mounted | Use it for |
|---------|-------|--------------|------------|
| `php` | `php:8.3-cli` | `php-track/php-app` | OOP, SOLID, GoF, pure-PHP DDD |
| `laravel` | `php:8.3-cli` + Laravel | `php-track/laravel-app` | architecture + DDD/CQRS in a real framework |
| `golang` | `golang:1.26` | `go-track/go-app` | everything in the Go track |

### First-time setup

```bash
git clone <this repo> && cd oop-patterns-learning
docker compose build          # builds all three images
```

### Running a lesson's practice

```bash
# Plain PHP — run a single example script:
docker compose run --rm php php examples/01-intro.php

# Go — run a lesson's demo command:
docker compose run --rm golang go run ./cmd/01-intro

# Laravel — boot the app, then hit a route in your browser:
docker compose up laravel     # first boot installs Composer deps
#   → open http://localhost:8000
```

> Prefer your local toolchain? You can also just `cd php-track/php-app && php examples/01-intro.php`
> or `cd go-track/go-app && go run ./cmd/01-intro` if you have PHP 8.3 / Go 1.26 installed.
> The Docker setup exists so the lessons have one command that always works.

---

## How a lesson is structured

Every `NN-topic.md` follows the same shape so you can build a rhythm:

1. **Goal** — what you'll be able to do afterward.
2. **Theory** — the concept, when to reach for it, the trade-offs, the common misuse.
3. **Language notes** — how PHP and Go differ on this specific topic.
4. **Practice** — exact files to open and the one command to run them.
5. **Exercises** — a few "now try…" prompts to make it stick.
6. **Recap & next** — links to the previous and next lesson.

---

## Course map (20 lessons per track)

```
Part 1 · OOP foundations          01 intro · 02 encapsulation · 03 abstraction
                                  04 inheritance-vs-composition · 05 polymorphism
Part 2 · SOLID                    06 SRP+OCP · 07 LSP+ISP · 08 DIP + DI
Part 3 · GoF patterns             09 creational · 10 structural · 11 behavioral
Part 4 · Architecture             12 layered · 13 hexagonal · 14 clean/onion
Part 5 · DDD + CQRS               15 strategic · 16 entities+VOs · 17 aggregates+events
                                  18 application layer · 19 CQRS + event sourcing
Part 6 · Capstone                 20 one bounded context, end to end
```

A small recurring domain — **ShopKit** (orders, catalog, pricing) — threads through the
pattern and architecture lessons so you watch isolated patterns compose into a real system.

---

## License

Personal learning material. Use it however helps you learn.
