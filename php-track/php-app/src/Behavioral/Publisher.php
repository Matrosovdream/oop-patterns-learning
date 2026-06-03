<?php

declare(strict_types=1);

namespace App\Behavioral;

/**
 * Observer — a subject notifies registered observers without knowing who they
 * are. In PHP, observers are most naturally plain callables.
 */
final class Publisher
{
    /** @var list<callable(Event): void> */
    private array $subscribers = [];

    public function subscribe(callable $subscriber): void
    {
        $this->subscribers[] = $subscriber;
    }

    public function publish(Event $event): void
    {
        foreach ($this->subscribers as $subscriber) {
            $subscriber($event);
        }
    }
}
