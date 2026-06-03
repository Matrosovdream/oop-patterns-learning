<?php

declare(strict_types=1);

namespace App\Creational;

use App\Money\Money;

final class CardPayment implements Payment
{
    public function charge(Money $amount): string
    {
        return "charged {$amount} to card";
    }
}
