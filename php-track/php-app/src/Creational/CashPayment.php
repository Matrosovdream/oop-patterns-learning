<?php

declare(strict_types=1);

namespace App\Creational;

use App\Money\Money;

final class CashPayment implements Payment
{
    public function charge(Money $amount): string
    {
        return "collected {$amount} in cash";
    }
}
