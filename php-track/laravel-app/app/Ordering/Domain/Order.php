<?php

declare(strict_types=1);

namespace App\Ordering\Domain;

use App\Ordering\Domain\Events\DomainEvent;
use App\Ordering\Domain\Events\OrderPaid;
use App\Ordering\Domain\Events\OrderPlaced;
use App\Ordering\Domain\Events\OrderShipped;
use App\Shared\Money;

/**
 * Order is the AGGREGATE ROOT of the Ordering context. Outside code may only
 * change it through these methods, which is what keeps its invariants ("a placed
 * order is immutable", "you can't ship before paying") always true. The state is
 * private precisely so nothing can bypass the rules.
 *
 * This is the same Order you'll wire through use cases (lesson 18) and persist
 * through a repository port (lesson 17) — the heart of the bounded context.
 */
final class Order
{
    private OrderStatus $status = OrderStatus::Draft;

    /** @var list<Line> */
    private array $lines = [];

    /** @var list<DomainEvent> */
    private array $events = [];

    public function __construct(
        private readonly OrderId $id,
        private readonly string $currency,
    ) {
    }

    public function addLine(Line $line): void
    {
        $this->guardDraft();
        $this->lines[] = $line;
    }

    public function place(): void
    {
        $this->guardDraft();
        if ($this->lines === []) {
            throw OrderException::emptyOrder();
        }
        $this->status = OrderStatus::Pending;
        $this->record(new OrderPlaced((string) $this->id, $this->total()));
    }

    public function pay(): void
    {
        if ($this->status !== OrderStatus::Pending) {
            throw OrderException::notPayable();
        }
        $this->status = OrderStatus::Paid;
        $this->record(new OrderPaid((string) $this->id, $this->total()));
    }

    public function ship(): void
    {
        if ($this->status !== OrderStatus::Paid) {
            throw OrderException::notShippable();
        }
        $this->status = OrderStatus::Shipped;
        $this->record(new OrderShipped((string) $this->id));
    }

    public function total(): Money
    {
        $total = Money::of(0, $this->currency);
        foreach ($this->lines as $line) {
            $total = $total->add($line->subtotal());
        }

        return $total;
    }

    public function id(): OrderId
    {
        return $this->id;
    }

    public function status(): OrderStatus
    {
        return $this->status;
    }

    /** @return list<Line> */
    public function lines(): array
    {
        return $this->lines;
    }

    /**
     * Returns events recorded since the last pull and clears them. The
     * application layer pulls these after saving and hands them to a publisher.
     *
     * @return list<DomainEvent>
     */
    public function pullEvents(): array
    {
        $events = $this->events;
        $this->events = [];

        return $events;
    }

    private function record(DomainEvent $event): void
    {
        $this->events[] = $event;
    }

    private function guardDraft(): void
    {
        if ($this->status !== OrderStatus::Draft) {
            throw OrderException::alreadyPlaced();
        }
    }
}
