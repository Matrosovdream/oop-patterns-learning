<?php

declare(strict_types=1);

namespace App\Banking;

use App\Money\Money;

/**
 * Vault can only RECEIVE money. It implements Depositable but NOT Withdrawable,
 * so the type system forbids passing it where a withdrawal is required — the
 * mistake is a type error, not a runtime "operation not supported" surprise.
 */
final class Vault implements Depositable
{
    public function __construct(private Money $balance)
    {
    }

    public function balance(): Money
    {
        return $this->balance;
    }

    public function deposit(Money $amount): void
    {
        $this->balance = $this->balance->add($amount);
    }
}
