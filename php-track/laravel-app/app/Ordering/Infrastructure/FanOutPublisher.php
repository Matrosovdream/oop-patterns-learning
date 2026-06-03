<?php

declare(strict_types=1);

namespace App\Ordering\Infrastructure;

use App\Ordering\Application\EventPublisher;

/**
 * Forwards events to several publishers — so the command side can drive a logger,
 * a recorder, AND the CQRS read-model projection from one event stream. It's the
 * Composite pattern (lesson 10) applied to a port: a group of publishers that is
 * itself a publisher.
 */
final class FanOutPublisher implements EventPublisher
{
    /** @var list<EventPublisher> */
    private array $targets;

    public function __construct(EventPublisher ...$targets)
    {
        $this->targets = $targets;
    }

    public function publish(array $events): void
    {
        foreach ($this->targets as $target) {
            $target->publish($events);
        }
    }
}
