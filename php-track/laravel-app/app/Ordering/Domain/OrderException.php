<?php

declare(strict_types=1);

namespace App\Ordering\Domain;

use DomainException;

/**
 * One exception type for the Ordering domain's invariant violations, with named
 * constructors so the message lives in the domain's vocabulary (not scattered
 * string literals). Application code can catch OrderException to map domain
 * failures to HTTP responses, etc.
 */
final class OrderException extends DomainException
{
    public static function emptyOrder(): self
    {
        return new self('cannot place an order with no lines');
    }

    public static function alreadyPlaced(): self
    {
        return new self('cannot change an order once it is placed');
    }

    public static function notPayable(): self
    {
        return new self('order cannot be paid from its current state');
    }

    public static function notShippable(): self
    {
        return new self('order cannot be shipped from its current state');
    }

    public static function notFound(string $id): self
    {
        return new self("order not found: {$id}");
    }
}
