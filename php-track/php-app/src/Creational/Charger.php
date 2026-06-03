<?php

declare(strict_types=1);

namespace App\Creational;

use App\Money\Money;

interface Charger
{
    public function charge(Money $amount): string;
}
