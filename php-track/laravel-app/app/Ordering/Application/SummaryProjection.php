<?php

declare(strict_types=1);

namespace App\Ordering\Application;

use App\Ordering\Domain\Events\OrderPaid;
use App\Ordering\Domain\Events\OrderPlaced;
use App\Ordering\Domain\Events\OrderShipped;
use App\Ordering\Domain\OrderStatus;

/**
 * The READ side of CQRS (lesson 19). It maintains denormalized OrderSummaries by
 * PROJECTING the same domain events the command side emits — so it implements the
 * EventPublisher port and can be plugged in right alongside a logger via the
 * FanOutPublisher. Queries never touch an Order aggregate.
 */
final class SummaryProjection implements EventPublisher
{
    /** @var array<string, OrderSummary> */
    private array $summaries = [];

    public function publish(array $events): void
    {
        foreach ($events as $event) {
            match (true) {
                $event instanceof OrderPlaced => $this->summaries[$event->orderId] =
                    new OrderSummary($event->orderId, OrderStatus::Pending->value, (string) $event->total),
                $event instanceof OrderPaid => $this->setStatus($event->orderId, OrderStatus::Paid),
                $event instanceof OrderShipped => $this->setStatus($event->orderId, OrderStatus::Shipped),
                default => null,
            };
        }
    }

    public function get(string $id): ?OrderSummary
    {
        return $this->summaries[$id] ?? null;
    }

    /** @return list<OrderSummary> */
    public function all(): array
    {
        $summaries = array_values($this->summaries);
        usort($summaries, static fn (OrderSummary $a, OrderSummary $b) => $a->id <=> $b->id);

        return $summaries;
    }

    private function setStatus(string $id, OrderStatus $status): void
    {
        if (! isset($this->summaries[$id])) {
            return;
        }
        $current = $this->summaries[$id];
        $this->summaries[$id] = new OrderSummary($current->id, $status->value, $current->total);
    }
}
