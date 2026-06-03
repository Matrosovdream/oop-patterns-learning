<?php

declare(strict_types=1);

namespace App\Banking;

use App\Money\Money;

/** A segregated capability (ISP): "money can be put in." */
interface Depositable
{
    public function deposit(Money $amount): void;
}
