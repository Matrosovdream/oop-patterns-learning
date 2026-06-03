<?php

declare(strict_types=1);

namespace App\Structural;

/** Simulates an expensive backend and counts how often it's actually called. */
final class SlowRates implements RateProvider
{
    public int $calls = 0;

    public function rate(string $currency): int
    {
        $this->calls++;

        return 108; // 100% + 8% tax, for demo purposes
    }
}
