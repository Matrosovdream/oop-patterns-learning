<?php

declare(strict_types=1);

namespace App\Notify;

final class SmsSender implements Sender
{
    public function __construct(private readonly string $gateway)
    {
    }

    public function send(string $to, string $message): void
    {
        printf("  sms via %s to %s: %s\n", $this->gateway, $to, $message);
    }
}
