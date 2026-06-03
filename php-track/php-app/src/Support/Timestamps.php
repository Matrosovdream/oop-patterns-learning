<?php

declare(strict_types=1);

namespace App\Support;

use DateTimeImmutable;

/**
 * Timestamps is a TRAIT — PHP's horizontal code-reuse mechanism. A class that
 * `use`s it gains these properties and methods as if they were declared inline.
 *
 * Traits are the closest PHP analog to Go's struct embedding: reuse WITHOUT an
 * inheritance "is-a" relationship. Use a trait for genuinely cross-cutting bits
 * of state+behavior (timestamps, soft-deletes) — not as a substitute for a
 * collaborator you should be composing (lesson 04 draws the line).
 */
trait Timestamps
{
    private ?DateTimeImmutable $createdAt = null;
    private ?DateTimeImmutable $updatedAt = null;

    public function touch(DateTimeImmutable $now): void
    {
        $this->createdAt ??= $now;
        $this->updatedAt = $now;
    }

    public function createdAt(): ?DateTimeImmutable
    {
        return $this->createdAt;
    }

    public function updatedAt(): ?DateTimeImmutable
    {
        return $this->updatedAt;
    }
}
