<?php

declare(strict_types=1);

namespace App\Creational;

use App\Money\Money;

interface Refunder
{
    public function refund(Money $amount): string;
}
