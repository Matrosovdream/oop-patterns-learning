<?php

declare(strict_types=1);

namespace App\Ordering\Application;

use App\Ordering\Domain\Line;
use App\Ordering\Domain\Order;
use App\Ordering\Domain\OrderId;
use App\Ordering\Domain\OrderRepository;
use App\Shared\Money;

/**
 * A USE CASE (application service). It orchestrates the domain and the ports —
 * load/build, call domain methods, save, publish — and holds NO business rules
 * itself (those live in the Order aggregate). Laravel's container injects the
 * OrderRepository and EventPublisher by type-hint (lesson 14).
 */
final class PlaceOrderHandler
{
    public function __construct(
        private readonly OrderRepository $orders,
        private readonly EventPublisher $publisher,
    ) {
    }

    public function handle(PlaceOrderCommand $command): OrderId
    {
        $order = new Order(new OrderId($command->orderId), $command->currency);

        foreach ($command->lines as $line) {
            $order->addLine(new Line(
                sku: $line->sku,
                name: $line->name,
                unit: Money::of($line->unitMinor, $command->currency),
                quantity: $line->quantity,
            ));
        }

        $order->place();
        $this->orders->save($order);
        $this->publisher->publish($order->pullEvents());

        return $order->id();
    }
}
