<?php

declare(strict_types=1);

namespace App\Ordering\Domain\Events;

/**
 * A domain event: a record that something meaningful happened, named in the past
 * tense and the ubiquitous language. The aggregate records these as it changes;
 * the application pulls and publishes them after saving.
 */
interface DomainEvent
{
    public function name(): string;
}
