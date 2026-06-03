<?php

declare(strict_types=1);

namespace App\Catalog;

/** A concrete ProductRepository backed by an array. A PDO/Eloquent one would
 *  implement the same interface. */
final class InMemoryProductRepository implements ProductRepository
{
    /** @var array<string, Product> */
    private array $items = [];

    public function save(Product $product): void
    {
        $this->items[$product->sku] = $product;
    }

    public function find(string $sku): ?Product
    {
        return $this->items[$sku] ?? null;
    }

    /** @return list<Product> */
    public function all(): array
    {
        return array_values($this->items);
    }
}
