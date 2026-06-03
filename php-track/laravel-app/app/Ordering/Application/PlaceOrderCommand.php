<?php

declare(strict_types=1);

namespace App\Ordering\Application;

/** The intent "place this order", as plain data. */
final class PlaceOrderCommand
{
    /** @param list<PlaceOrderLine> $lines */
    public function __construct(
        public readonly string $orderId,
        public readonly string $currency,
        public readonly array $lines,
    ) {
    }
}
