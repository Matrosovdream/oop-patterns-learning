<?php

declare(strict_types=1);

namespace App\Ordering\Domain\Events;

final class OrderShipped implements DomainEvent
{
    public function __construct(public readonly string $orderId)
    {
    }

    public function name(): string
    {
        return 'order.shipped';
    }
}
