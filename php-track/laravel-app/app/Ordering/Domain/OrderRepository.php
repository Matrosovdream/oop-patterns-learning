<?php

declare(strict_types=1);

namespace App\Ordering\Domain;

/**
 * OrderRepository is a PORT: an interface the domain defines in its own terms
 * (save/find a whole Order by identity), leaving the "how" (in-memory, Eloquent,
 * an API) to an adapter in the Infrastructure layer. The interface lives HERE, in
 * the domain — the database adapter depends on this, not the reverse. That
 * inverted dependency is what the hexagonal architecture (lesson 13) is built on.
 *
 * A repository deals in whole aggregates, never rows: find() returns a fully
 * reconstituted Order with its invariants intact.
 */
interface OrderRepository
{
    public function save(Order $order): void;

    /** @throws OrderException when no order has the given id. */
    public function find(OrderId $id): Order;
}
