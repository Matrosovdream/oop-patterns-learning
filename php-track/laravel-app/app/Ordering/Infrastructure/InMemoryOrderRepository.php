<?php

declare(strict_types=1);

namespace App\Ordering\Infrastructure;

use App\Ordering\Domain\Order;
use App\Ordering\Domain\OrderException;
use App\Ordering\Domain\OrderId;
use App\Ordering\Domain\OrderRepository;

/**
 * A driven ADAPTER implementing the OrderRepository port with an array. An
 * Eloquent adapter would implement the SAME interface (map the aggregate to/from
 * rows) and the application layer wouldn't notice the swap — that's the point of
 * the port. Bound as a singleton in OrderingServiceProvider so state survives
 * across a single request/command for the demos.
 */
final class InMemoryOrderRepository implements OrderRepository
{
    /** @var array<string, Order> */
    private array $orders = [];

    public function save(Order $order): void
    {
        $this->orders[(string) $order->id()] = $order;
    }

    public function find(OrderId $id): Order
    {
        return $this->orders[(string) $id] ?? throw OrderException::notFound((string) $id);
    }
}
