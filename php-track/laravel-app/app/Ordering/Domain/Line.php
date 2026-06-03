<?php

declare(strict_types=1);

namespace App\Ordering\Domain;

use App\Shared\Money;

/**
 * Line is a value object living INSIDE the Order aggregate. It has no identity of
 * its own and is only ever reached through its Order.
 */
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
