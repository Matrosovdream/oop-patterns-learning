<?php

declare(strict_types=1);

namespace App\Behavioral;

final class Event
{
    public function __construct(
        public readonly string $name,
        public readonly string $orderId,
    ) {
    }
}
