<?php

declare(strict_types=1);

namespace App\Creational;

use InvalidArgumentException;

final class GatewayFactory
{
    public function make(string $name): Gateway
    {
        return match ($name) {
            'stripe' => new StripeGateway(),
            default => throw new InvalidArgumentException("unknown gateway: {$name}"),
        };
    }
}
