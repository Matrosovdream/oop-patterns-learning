<?php

declare(strict_types=1);

// The "does my environment work?" lesson.
//   docker compose run --rm php php examples/01-intro.php
//   # or locally: cd php-track/php-app && composer install -q && php examples/01-intro.php

require __DIR__ . '/../vendor/autoload.php';

use App\Money\Money;

echo "PHP track is alive. Welcome to OOP-by-design.\n\n";

$price = Money::fromMajor(19, 99, 'USD');
$withTax = $price->percentage(108); // 108% == price + 8% tax

printf("list price:   %s\n", $price);
printf("with 8%% tax:  %s\n", $withTax);

echo "\nNext: lesson 02 — classes, objects & encapsulation.\n";
