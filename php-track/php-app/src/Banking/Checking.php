<?php

declare(strict_types=1);

namespace App\Banking;

use App\Money\Money;

/**
 * Checking supports both deposits and withdrawals, so it implements BOTH
 * capabilities — and it honors the Withdrawable contract fully (LSP): a
 * withdrawal either succeeds or throws the documented exception.
 */
final class Checking implements Depositable, Withdrawable
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

    public function withdraw(Money $amount): void
    {
        if ($this->balance->amount() < $amount->amount()) {
            throw new InsufficientFundsException('insufficient funds');
        }
        $this->balance = $this->balance->subtract($amount);
    }
}
