<?php

declare(strict_types=1);

namespace App\Ordering\Domain;

use Stringable;

/** The aggregate's identity — a value object wrapping the raw id string. */
final class OrderId implements Stringable
{
    public function __construct(public readonly string $value)
    {
    }

    public function equals(OrderId $other): bool
    {
        return $this->value === $other->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}
