<?php

declare(strict_types=1);

namespace App\Behavioral;

final class PendingState implements OrderState
{
    public function pay(): OrderState
    {
        return new PaidState();
    }

    public function ship(): OrderState
    {
        throw new IllegalTransitionException('cannot ship an unpaid order');
    }

    public function name(): string
    {
        return 'pending';
    }
}
