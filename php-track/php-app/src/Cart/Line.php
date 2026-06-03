<?php

declare(strict_types=1);

namespace App\Cart;

use App\Money\Money;

/** A single row in the cart — an immutable value object. */
final class Line
{
    public function __construct(
        public readonly string $sku,
        public readonly string $name,
        public readonly Money $unit,
        public readonly int $quantity,
    ) {
    }

    public function subtotal(): Money
    {
        return $this->unit->multiply($this->quantity);
    }
}
