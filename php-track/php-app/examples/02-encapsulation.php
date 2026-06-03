<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use App\Money\Money;

// Money can only be built through a named constructor — the constructor is
// private and the properties are readonly, so nothing can put it into an invalid
// state from outside.
$wallet = Money::fromMajor(50, 0, 'USD');
$coffee = Money::fromMajor(4, 75, 'USD');

$remaining = $wallet->subtract($coffee);
printf("wallet %s - coffee %s = %s\n", $wallet, $coffee, $remaining);

// $wallet is unchanged: Money is immutable, so subtract() returned a NEW value.
printf("wallet is still %s (immutability)\n", $wallet);

// Encapsulation lets the type reject nonsense — mixing currencies is a domain error.
try {
    $wallet->add(Money::fromMajor(10, 0, 'EUR'));
} catch (InvalidArgumentException $e) {
    printf("blocked bad operation: %s\n", $e->getMessage());
}
