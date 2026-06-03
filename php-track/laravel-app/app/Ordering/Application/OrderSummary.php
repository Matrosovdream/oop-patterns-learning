<?php

declare(strict_types=1);

namespace App\Ordering\Application;

/** A flat, display-ready view for the CQRS read side — no behavior, no invariants. */
final class OrderSummary
{
    public function __construct(
        public readonly string $id,
        public readonly string $status,
        public readonly string $total,
    ) {
    }

    /** @return array{id: string, status: string, total: string} */
    public function toArray(): array
    {
        return ['id' => $this->id, 'status' => $this->status, 'total' => $this->total];
    }
}
