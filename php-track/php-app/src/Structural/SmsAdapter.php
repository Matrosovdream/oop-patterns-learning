<?php

declare(strict_types=1);

namespace App\Structural;

use App\Notify\Sender;

/**
 * Adapter — make the incompatible LegacySms fit the Sender interface our code
 * already speaks, so legacy code drops into anything expecting a Sender.
 */
final class SmsAdapter implements Sender
{
    public function __construct(private readonly LegacySms $client)
    {
    }

    public function send(string $to, string $message): void
    {
        $this->client->sendText($to, $message);
    }
}
