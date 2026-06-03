<?php

declare(strict_types=1);

namespace App\Cart;

use App\Money\Money;

/**
 * Cart has ONE responsibility: hold lines and total them. It knows nothing about
 * discounts, tax, persistence, or display — those are other classes' jobs
 * (lesson 06's Single Responsibility Principle).
 */
final class Cart
{
    /** @var list<Line> */
    private array $lines = [];

    public function __construct(private readonly string $currency)
    {
    }

    public function add(Line $line): void
    {
        $this->lines[] = $line;
    }

    /** @return list<Line> */
    public function lines(): array
    {
        return $this->lines;
    }

    public function currency(): string
    {
        return $this->currency;
    }

    public function subtotal(): Money
    {
        $total = Money::of(0, $this->currency);
        foreach ($this->lines as $line) {
            $total = $total->add($line->subtotal());
        }

        return $total;
    }
}
