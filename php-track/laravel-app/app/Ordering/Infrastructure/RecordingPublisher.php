<?php

declare(strict_types=1);

namespace App\Ordering\Infrastructure;

use App\Ordering\Application\EventPublisher;

/**
 * Another adapter for the EventPublisher port — it just captures event names in
 * memory. It doubles as a test double (lesson 13: swap an adapter without
 * touching the use case) and as the demo's event display.
 */
final class RecordingPublisher implements EventPublisher
{
    /** @var list<string> */
    private array $names = [];

    public function publish(array $events): void
    {
        foreach ($events as $event) {
            $this->names[] = $event->name();
        }
    }

    /** @return list<string> */
    public function names(): array
    {
        return $this->names;
    }
}
