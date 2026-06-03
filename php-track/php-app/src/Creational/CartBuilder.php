<?php

declare(strict_types=1);

namespace App\Creational;

use App\Cart\Cart;
use App\Cart\Line;
use App\Money\Money;

/**
 * Builder — assemble a complex object step by step with a fluent API. Each
 * `add()` returns `$this` so calls chain; `build()` produces the finished Cart.
 */
final class CartBuilder
{
    /** @var list<Line> */
    private array $lines = [];

    public function __construct(private readonly string $currency)
    {
    }

    public function add(string $sku, string $name, Money $unit, int $quantity): self
    {
        $this->lines[] = new Line($sku, $name, $unit, $quantity);

        return $this;
    }

    public function build(): Cart
    {
        $cart = new Cart($this->currency);
        foreach ($this->lines as $line) {
            $cart->add($line);
        }

        return $cart;
    }
}
