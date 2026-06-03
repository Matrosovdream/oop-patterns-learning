<?php

declare(strict_types=1);

namespace App\Structural;

use App\Notify\Sender;

/**
 * Decorator — wrap a Sender to add behavior (here, a tag) while keeping the SAME
 * interface. Because it's a Sender too, decorators stack: new PrefixSender(new
 * RetrySender(new EmailSender(...))).
 */
final class PrefixSender implements Sender
{
    public function __construct(
        private readonly Sender $inner,
        private readonly string $prefix,
    ) {
    }

    public function send(string $to, string $message): void
    {
        $this->inner->send($to, $this->prefix . $message);
    }
}
