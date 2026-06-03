<?php

declare(strict_types=1);

namespace App\Ordering\Application;

use App\Ordering\Domain\Events\DomainEvent;

/**
 * An OUTPUT PORT. The application hands domain events to "the outside" without
 * knowing whether they're logged, queued, or projected into a read model.
 * Infrastructure adapters (lesson 13) implement it; the service container binds
 * the concrete one (lesson 14).
 */
interface EventPublisher
{
    /** @param list<DomainEvent> $events */
    public function publish(array $events): void;
}
