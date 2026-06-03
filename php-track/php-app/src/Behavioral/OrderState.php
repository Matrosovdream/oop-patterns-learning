<?php

declare(strict_types=1);

namespace App\Behavioral;

/**
 * State — when behavior depends on state, model each state as a type that knows
 * which transitions are legal. Transitions return the next state, or throw if the
 * move isn't allowed from here. This replaces a brittle `if ($status === ...)`
 * sprawl with an explicit state machine.
 */
interface OrderState
{
    public function pay(): OrderState;

    public function ship(): OrderState;

    public function name(): string;
}
