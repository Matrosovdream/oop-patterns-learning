<?php

declare(strict_types=1);

namespace App\Behavioral;

final class ShippedState implements OrderState
{
    public function pay(): OrderState
    {
        throw new IllegalTransitionException('a shipped order cannot be paid again');
    }

    public function ship(): OrderState
    {
        throw new IllegalTransitionException('order is already shipped');
    }

    public function name(): string
    {
        return 'shipped';
    }
}
