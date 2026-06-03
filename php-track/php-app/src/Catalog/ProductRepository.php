<?php

declare(strict_types=1);

namespace App\Catalog;

/**
 * ProductRepository is the abstraction the service depends on for persistence —
 * expressed in domain terms (save/find/all a Product), not SQL. Concrete stores
 * implement it. Defining the interface here, beside what uses it, is the seam
 * that lets the hexagonal architecture (lesson 13) swap adapters freely.
 */
interface ProductRepository
{
    public function save(Product $product): void;

    public function find(string $sku): ?Product;

    /** @return list<Product> */
    public function all(): array;
}
