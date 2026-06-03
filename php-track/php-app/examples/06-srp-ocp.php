<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use App\Cart\Cart;
use App\Cart\Line;
use App\Money\Money;
use App\Notify\EmailSender;
use App\Pricing\Discount;
use App\Pricing\PercentageOff;

// SRP: four responsibilities, four separate places — the cart totals, the
// discount decides a reduction, renderReceipt() formats, the sender delivers.
// OCP: checkout() takes a Discount, so a new discount type never edits it.

function checkout(Cart $cart, Discount $discount): array
{
    $subtotal = $cart->subtotal();
    $total = $discount->apply($subtotal);

    return [$subtotal, $total];
}

function renderReceipt(Cart $cart, Discount $discount, Money $subtotal, Money $total): string
{
    $out = "ShopKit receipt\n";
    foreach ($cart->lines() as $line) {
        $out .= sprintf("  %-14s %d × %s = %s\n", $line->name, $line->quantity, $line->unit, $line->subtotal());
    }
    $out .= sprintf("  subtotal:     %s\n", $subtotal);
    $out .= sprintf("  discount:     %s\n", $discount->label());
    $out .= sprintf("  total:        %s\n", $total);

    return $out;
}

$cart = new Cart('USD');
$cart->add(new Line('MUG-1', 'Enamel Mug', Money::fromMajor(12, 0, 'USD'), 2));
$cart->add(new Line('TEE-1', 'Logo Tee', Money::fromMajor(25, 0, 'USD'), 1));

$discount = new PercentageOff(10);          // OCP: swap freely, checkout() unchanged
[$subtotal, $total] = checkout($cart, $discount);

$receipt = renderReceipt($cart, $discount, $subtotal, $total);  // SRP: format only
(new EmailSender('orders@shopkit.test'))->send('buyer@example.com', "\n" . $receipt); // SRP: deliver only
