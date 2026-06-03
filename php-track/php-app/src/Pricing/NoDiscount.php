<?php

declare(strict_types=1);

namespace App\Pricing;

use App\Money\Money;

/** Leaves the price untouched — a useful Null Object (lesson 10). */
final class NoDiscount implements Discount
{
    public function apply(Money $price): Money
    {
        return $price;
    }

    public function label(): string
    {
        return 'no discount';
    }
}
