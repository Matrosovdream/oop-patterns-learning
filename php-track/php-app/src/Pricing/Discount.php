<?php

declare(strict_types=1);

namespace App\Pricing;

use App\Money\Money;

/**
 * Discount is the abstraction every discount rule implements. It's the seed of
 * the Strategy pattern (lesson 11): interchangeable behaviors behind one type,
 * chosen at runtime.
 */
interface Discount
{
    public function apply(Money $price): Money;

    public function label(): string;
}
