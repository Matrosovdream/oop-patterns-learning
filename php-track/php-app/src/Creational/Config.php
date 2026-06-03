<?php

declare(strict_types=1);

namespace App\Creational;

/**
 * Singleton — exactly one instance, reached through a global access point. The
 * private constructor blocks `new Config()` from outside.
 *
 * Use sparingly: a singleton is a global, and globals fight dependency injection
 * and make testing harder. In a Laravel app you'd register a single instance in
 * the container instead (a "scoped"/singleton binding) and inject it — same "one
 * instance", but without the global.
 */
final class Config
{
    private static ?self $instance = null;

    private function __construct(
        public readonly string $currency,
        public readonly int $taxPercent,
    ) {
    }

    public static function instance(): self
    {
        return self::$instance ??= new self(currency: 'USD', taxPercent: 8);
    }
}
