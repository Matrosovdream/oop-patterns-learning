<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use App\Creational\CartBuilder;
use App\Creational\Config;
use App\Creational\GatewayFactory;
use App\Creational\PaymentFactory;
use App\Money\Money;

$price = Money::fromMajor(42, 0, 'USD');

// Factory Method: ask for a kind, get a Payment; concrete class stays hidden.
$payments = new PaymentFactory();
foreach (['card', 'cash'] as $kind) {
    printf("[factory]   %s\n", $payments->make($kind)->charge($price));
}

// Abstract Factory: one family yields a matching charger + refunder.
$gateway = (new GatewayFactory())->make('stripe');
printf("[abstract]  %s; %s\n", $gateway->charger()->charge($price), $gateway->refunder()->refund($price));

// Builder: assemble a cart fluently, then build the finished value.
$cart = (new CartBuilder('USD'))
    ->add('MUG-1', 'Enamel Mug', Money::fromMajor(12, 0, 'USD'), 2)
    ->add('TEE-1', 'Logo Tee', Money::fromMajor(25, 0, 'USD'), 1)
    ->build();
printf("[builder]   cart subtotal %s\n", $cart->subtotal());

// Singleton: same instance every call.
$a = Config::instance();
$b = Config::instance();
printf("[singleton] same instance? %s (currency=%s, tax=%d%%)\n", $a === $b ? 'true' : 'false', $a->currency, $a->taxPercent);
