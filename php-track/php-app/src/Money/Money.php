<?php

declare(strict_types=1);

namespace App\Money;

use InvalidArgumentException;

/**
 * Money is a small, immutable value object for amounts in a single currency.
 *
 * Two design choices to notice:
 *  - Amounts are stored as integer minor units (cents), never float — money and
 *    floating point do not mix (0.1 + 0.2 !== 0.3).
 *  - The class is `final`, the constructor is `private`, and the properties are
 *    `readonly`. The only way to get a Money is via a named constructor, and the
 *    only way to "change" one is to derive a new one. That is encapsulation
 *    protecting an invariant.
 */
final class Money
{
    private function __construct(
        private readonly int $minor,
        private readonly string $currency,
    ) {
    }

    /** Build from a raw minor-unit amount (cents). */
    public static function of(int $minorUnits, string $currency): self
    {
        return new self($minorUnits, $currency);
    }

    /** Build from major + minor units, e.g. fromMajor(19, 99, 'USD') === $19.99. */
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

    public function isZero(): bool
    {
        return $this->minor === 0;
    }

    /** Returns a NEW Money; the receiver is never mutated. */
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

    /** Scale the amount, e.g. unit price × quantity. */
    public function multiply(int $factor): self
    {
        return new self($this->minor * $factor, $this->currency);
    }

    /** Whole-number percentage of the amount, rounded to the nearest cent. */
    public function percentage(int $percent): self
    {
        return new self(intdiv($this->minor * $percent + 50, 100), $this->currency);
    }

    /** Value equality: equal contents, not the same object. */
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
            throw new InvalidArgumentException(
                "currency mismatch: {$this->currency} vs {$other->currency}"
            );
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
