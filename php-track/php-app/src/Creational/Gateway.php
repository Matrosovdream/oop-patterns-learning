<?php

declare(strict_types=1);

namespace App\Creational;

/**
 * Abstract Factory — each concrete gateway produces a FAMILY of related objects
 * (a charger + a matching refunder) that belong together, so you can't pair
 * Stripe's charger with PayPal's refunder by accident.
 */
interface Gateway
{
    public function name(): string;

    public function charger(): Charger;

    public function refunder(): Refunder;
}
