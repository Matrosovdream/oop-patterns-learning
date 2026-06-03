<?php

declare(strict_types=1);

namespace App\Ordering\Domain\Events;

use App\Shared\Money;

final class OrderPlaced implements DomainEvent
{
    public function __construct(
        public readonly string $orderId,
        public readonly Money $total,
    ) {
    }

    public function name(): string
    {
        return 'order.placed';
    }
}
