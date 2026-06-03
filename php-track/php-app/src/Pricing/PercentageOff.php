<?php

declare(strict_types=1);

namespace App\Pricing;

use App\Money\Money;

final class PercentageOff implements Discount
{
    public function __construct(private readonly int $percent)
    {
    }

    public function apply(Money $price): Money
    {
        return $price->subtract($price->percentage($this->percent));
    }

    public function label(): string
    {
        return "{$this->percent}% off";
    }
}
