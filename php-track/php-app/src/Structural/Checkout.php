<?php

declare(strict_types=1);

namespace App\Structural;

use App\Cart\Cart;
use App\Money\Money;
use App\Notify\Sender;
use App\Pricing\Discount;

/**
 * Facade — one simple entry point over several subsystems (cart total, pricing,
 * notification). Callers get a single placeOrder() instead of orchestrating three
 * collaborators themselves.
 */
final class Checkout
{
    public function __construct(private readonly Sender $notifier)
    {
    }

    public function placeOrder(Cart $cart, Discount $discount, string $email): Money
    {
        $total = $discount->apply($cart->subtotal());
        $this->notifier->send($email, "Thanks! Your order total is {$total}");

        return $total;
    }
}
