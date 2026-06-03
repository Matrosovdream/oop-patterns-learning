<?php

declare(strict_types=1);

namespace App\Ordering\Application;

/**
 * An input DTO — plain transport data, NOT a domain type. Keeping DTOs separate
 * from domain objects is the anti-corruption boundary (lesson 18): the outside
 * world's shapes (JSON, form fields) don't leak into the model. Note it carries a
 * primitive `unitMinor`, not a Money value object — the handler builds the VO.
 */
final class PlaceOrderLine
{
    public function __construct(
        public readonly string $sku,
        public readonly string $name,
        public readonly int $unitMinor,
        public readonly int $quantity,
    ) {
    }
}
