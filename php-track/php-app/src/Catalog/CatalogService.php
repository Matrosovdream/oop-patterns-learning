<?php

declare(strict_types=1);

namespace App\Catalog;

use App\Money\Money;
use App\Notify\Sender;
use App\Support\Clock;

/**
 * CatalogService is HIGH-LEVEL policy. The Dependency Inversion Principle says it
 * must depend on abstractions, not details — so it takes a ProductRepository, a
 * Sender, and a Clock (all interfaces) through its constructor. It names no
 * database, no mailer, and not even the system clock. Those concrete choices are
 * injected from the outside (see examples/08-dip-di.php).
 */
final class CatalogService
{
    public function __construct(
        private readonly ProductRepository $repository,
        private readonly Sender $notifier,
        private readonly Clock $clock,
    ) {
    }

    public function addProduct(string $sku, string $name, Money $price): Product
    {
        $product = Product::create($sku, $name, $price, $this->clock->now());
        $this->repository->save($product);
        $this->notifier->send('ops@shopkit.test', sprintf('new product: %s @ %s', $product->name, $product->price()));

        return $product;
    }

    public function count(): int
    {
        return count($this->repository->all());
    }
}
