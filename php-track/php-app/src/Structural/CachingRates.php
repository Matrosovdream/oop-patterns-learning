<?php

declare(strict_types=1);

namespace App\Structural;

/**
 * Proxy — same interface as the real RateProvider, but it controls access by
 * caching, so the slow backend is hit at most once per currency.
 *
 * (Decorator vs Proxy: identical shape. A decorator ADDS behavior the caller
 * wants; a proxy CONTROLS ACCESS to the real subject. Intent is the difference.)
 */
final class CachingRates implements RateProvider
{
    /** @var array<string, int> */
    private array $cache = [];

    public function __construct(private readonly RateProvider $inner)
    {
    }

    public function rate(string $currency): int
    {
        return $this->cache[$currency] ??= $this->inner->rate($currency);
    }
}
