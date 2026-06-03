<?php

declare(strict_types=1);

namespace App\Behavioral;

final class PaidState implements OrderState
{
    public function pay(): OrderState
    {
        throw new IllegalTransitionException('order is already paid');
    }

    public function ship(): OrderState
    {
        return new ShippedState();
    }

    public function name(): string
    {
        return 'paid';
    }
}
