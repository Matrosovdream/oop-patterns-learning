<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use App\Banking\Checking;
use App\Banking\Depositable;
use App\Banking\InsufficientFundsException;
use App\Banking\Vault;
use App\Banking\Withdrawable;
use App\Money\Money;

// transfer() asks for the two narrow capabilities it uses, not a fat "Account".
// Anything Withdrawable can be the source; anything Depositable can be the sink.
function transfer(Withdrawable $from, Depositable $to, Money $amount): void
{
    $from->withdraw($amount);
    $to->deposit($amount);
}

$checking = new Checking(Money::fromMajor(100, 0, 'USD'));
$vault = new Vault(Money::fromMajor(0, 0, 'USD'));

// Checking is Withdrawable AND Depositable, so this is fine.
transfer($checking, $vault, Money::fromMajor(30, 0, 'USD'));
printf("checking: %s, vault: %s\n", $checking->balance(), $vault->balance());

// The LSP win: `transfer($vault, $checking, ...)` is a TYPE ERROR — Vault is not
// Withdrawable, so it can't be the source. Uncomment to see PHP reject it:
//   transfer($vault, $checking, Money::fromMajor(10, 0, 'USD'));
//   // TypeError: Argument #1 ($from) must be of type Withdrawable, Vault given

// Over-withdrawing throws the documented exception (contract honored, not a panic).
try {
    transfer($checking, $vault, Money::fromMajor(1000, 0, 'USD'));
} catch (InsufficientFundsException $e) {
    printf("blocked over-withdrawal: %s\n", $e->getMessage());
}
