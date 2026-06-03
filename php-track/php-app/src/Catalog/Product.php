<?php

declare(strict_types=1);

namespace App\Catalog;

use App\Money\Money;
use App\Support\Timestamps;
use DateTimeImmutable;

/**
 * Product is a ShopKit catalog item. It demonstrates composition two ways:
 *   - it USES the Timestamps trait (horizontal reuse, no inheritance), and
 *   - it HOLDS a Money price as a field.
 * Neither is inheritance — Product `extends` nothing.
 */
final class Product
{
    use Timestamps;

    public function __construct(
        public readonly string $sku,
        public readonly string $name,
        private Money $price,
    ) {
    }

    public static function create(string $sku, string $name, Money $price, DateTimeImmutable $now): self
    {
        $product = new self($sku, $name, $price);
        $product->touch($now);

        return $product;
    }

    public function price(): Money
    {
        return $this->price;
    }

    public function reprice(Money $newPrice, DateTimeImmutable $now): void
    {
        $this->price = $newPrice;
        $this->touch($now);
    }
}
