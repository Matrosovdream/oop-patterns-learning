<?php

declare(strict_types=1);

namespace App\Shared;

use InvalidArgumentException;

/**
 * Money is a Shared Kernel value object — a small immutable type shared by more
 * than one bounded context (Ordering, Pricing, …) by explicit agreement. It is
 * the same Money you met in the plain-PHP track, living here in App\Shared.
 *
 * Immutable (private readonly fields, derive-a-new-one methods), integer minor
 * units (never float), and currency-safe (refuses to mix currencies).
 */
final class Money
{
    private function __construct(
        private readonly int $minor,
        private readonly string $currency,
    ) {
    }

    public static function of(int $minorUnits, string $currency): self
    {
        return new self($minorUnits, $currency);
    }

    public static function fromMajor(int $units, int $cents, string $currency): self
    {
        return new self($units * 100 + $cents, $currency);
    }

    public function amount(): int
    {
        return $this->minor;
    }

    public function currency(): string
    {
        return $this->currency;
    }

    public function add(Money $other): self
    {
        $this->assertSameCurrency($other);

        return new self($this->minor + $other->minor, $this->currency);
    }

    public function subtract(Money $other): self
    {
        $this->assertSameCurrency($other);

        return new self($this->minor - $other->minor, $this->currency);
    }

    public function multiply(int $factor): self
    {
        return new self($this->minor * $factor, $this->currency);
    }

    public function equals(Money $other): bool
    {
        return $this->minor === $other->minor && $this->currency === $other->currency;
    }

    public function __toString(): string
    {
        $abs = abs($this->minor);
        $sign = $this->minor < 0 ? '-' : '';

        return sprintf('%s%s%d.%02d', $sign, self::symbol($this->currency), intdiv($abs, 100), $abs % 100);
    }

    private function assertSameCurrency(Money $other): void
    {
        if ($this->currency !== $other->currency) {
            throw new InvalidArgumentException("currency mismatch: {$this->currency} vs {$other->currency}");
        }
    }

    private static function symbol(string $currency): string
    {
        return match ($currency) {
            'USD' => '$',
            'EUR' => '€',
            'GBP' => '£',
            default => $currency . ' ',
        };
    }
}
