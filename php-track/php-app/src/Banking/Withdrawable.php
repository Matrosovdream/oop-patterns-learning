<?php

declare(strict_types=1);

namespace App\Banking;

use App\Money\Money;

/**
 * A segregated capability (ISP): "money can be taken out." Splitting this from
 * Depositable means a type only promises what it can actually honor — which is
 * what prevents Liskov violations (a Vault that "supports" withdraw but throws).
 */
interface Withdrawable
{
    /** @throws InsufficientFundsException when the balance is too low. */
    public function withdraw(Money $amount): void;
}
