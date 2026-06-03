<?php

declare(strict_types=1);

namespace App\Ordering\Infrastructure;

use App\Ordering\Application\EventPublisher;
use Illuminate\Support\Facades\Log;

/**
 * A driven adapter for the EventPublisher port that writes events to Laravel's
 * log. A production adapter might push to a queue, an outbox table, or a message
 * bus — all behind the same port.
 */
final class LoggingPublisher implements EventPublisher
{
    public function publish(array $events): void
    {
        foreach ($events as $event) {
            Log::info('ordering event', ['event' => $event->name()]);
        }
    }
}
