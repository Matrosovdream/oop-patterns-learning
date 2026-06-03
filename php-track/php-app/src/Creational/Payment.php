<?php

declare(strict_types=1);

namespace App\Creational;

use App\Money\Money;

/** The product type the factory hands back. */
interface Payment
{
    public function charge(Money $amount): string;
}
