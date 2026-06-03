<?php

declare(strict_types=1);

namespace App\Notify;

final class EmailSender implements Sender
{
    public function __construct(private readonly string $from)
    {
    }

    public function send(string $to, string $message): void
    {
        printf("  email from %s to %s: %s\n", $this->from, $to, $message);
    }
}
