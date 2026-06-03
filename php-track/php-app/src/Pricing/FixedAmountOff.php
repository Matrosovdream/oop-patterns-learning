<?php

declare(strict_types=1);

namespace App\Pricing;

use App\Money\Money;

final class FixedAmountOff implements Discount
{
    public function __construct(private readonly Money $amount)
    {
    }

    public function apply(Money $price): Money
    {
        $result = $price->subtract($this->amount);

        return $result->amount() < 0 ? Money::of(0, $price->currency()) : $result;
    }

    public function label(): string
    {
        return $this->amount . ' off';
    }
}
