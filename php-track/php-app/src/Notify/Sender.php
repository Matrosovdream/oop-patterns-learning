<?php

declare(strict_types=1);

namespace App\Notify;

/**
 * Sender is the abstraction: anything that can deliver a message. Callers depend
 * on this interface, never on a concrete email/SMS class — so the mechanism can
 * change without touching them. In PHP, implementers declare `implements Sender`
 * explicitly (unlike Go's implicit satisfaction).
 */
interface Sender
{
    public function send(string $to, string $message): void;
}
