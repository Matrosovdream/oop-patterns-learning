<?php

declare(strict_types=1);

namespace App\Notify;

/**
 * Broadcast works against the Sender abstraction, so it can deliver through ANY
 * implementation — today's email/SMS, tomorrow's Slack — without changing here.
 */
final class Broadcast
{
    /** @param list<string> $recipients */
    public static function to(Sender $sender, array $recipients, string $message): void
    {
        foreach ($recipients as $recipient) {
            $sender->send($recipient, $message);
        }
    }
}
