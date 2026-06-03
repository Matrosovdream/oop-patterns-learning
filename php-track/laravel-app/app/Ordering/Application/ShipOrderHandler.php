<?php

declare(strict_types=1);

namespace App\Ordering\Application;

use App\Ordering\Domain\OrderId;
use App\Ordering\Domain\OrderRepository;

final class ShipOrderHandler
{
    public function __construct(
        private readonly OrderRepository $orders,
        private readonly EventPublisher $publisher,
    ) {
    }

    public function handle(string $orderId): void
    {
        $order = $this->orders->find(new OrderId($orderId));
        $order->ship();
        $this->orders->save($order);
        $this->publisher->publish($order->pullEvents());
    }
}
