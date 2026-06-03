<?php

declare(strict_types=1);

namespace App\Creational;

use App\Money\Money;

/**
 * A concrete factory. It returns its family as small anonymous classes — an
 * idiomatic PHP way to provide one-off implementations without extra files.
 */
final class StripeGateway implements Gateway
{
    public function name(): string
    {
        return 'stripe';
    }

    public function charger(): Charger
    {
        return new class implements Charger {
            public function charge(Money $amount): string
            {
                return "stripe charge {$amount}";
            }
        };
    }

    public function refunder(): Refunder
    {
        return new class implements Refunder {
            public function refund(Money $amount): string
            {
                return "stripe refund {$amount}";
            }
        };
    }
}
