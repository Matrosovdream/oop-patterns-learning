<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use App\Money\Money;
use App\Pricing\Discount;
use App\Pricing\FixedAmountOff;
use App\Pricing\NoDiscount;
use App\Pricing\PercentageOff;

$listPrice = Money::fromMajor(80, 0, 'USD');

// Every element implements Discount, so we can hold them in one array and call
// apply()/label() without caring which concrete class each one is.
/** @var list<Discount> $discounts */
$discounts = [
    new NoDiscount(),
    new PercentageOff(10),
    new PercentageOff(25),
    new FixedAmountOff(Money::fromMajor(15, 0, 'USD')),
];

printf("list price: %s\n\n", $listPrice);
foreach ($discounts as $discount) {
    printf("%-12s -> %s\n", $discount->label(), $discount->apply($listPrice));
}

echo "\nOne loop, many behaviors — polymorphism through a shared interface.\n";
