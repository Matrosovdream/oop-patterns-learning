<?php

declare(strict_types=1);

namespace App\Creational;

use InvalidArgumentException;

/**
 * Factory Method — callers name WHAT they want ("card") and get a Payment, never
 * touching the concrete classes. Add a kind here and no caller changes.
 */
final class PaymentFactory
{
    public function make(string $kind): Payment
    {
        return match ($kind) {
            'card' => new CardPayment(),
            'cash' => new CashPayment(),
            default => throw new InvalidArgumentException("unknown payment kind: {$kind}"),
        };
    }
}
