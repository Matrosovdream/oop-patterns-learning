<?php

declare(strict_types=1);

namespace App\Structural;

/** A pretend third-party client we don't control. Its method shape (sendText)
 *  doesn't match our App\Notify\Sender (send). */
final class LegacySms
{
    public function sendText(string $number, string $body): void
    {
        printf("  [legacy-sms] to %s: %s\n", $number, $body);
    }
}
