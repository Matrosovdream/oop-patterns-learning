<?php

declare(strict_types=1);

namespace App\Behavioral;

final class Deposit implements Command
{
    public function __construct(
        private readonly Ledger $ledger,
        private readonly int $amount,
    ) {
    }

    public function execute(): void
    {
        $this->ledger->balance += $this->amount;
    }

    public function undo(): void
    {
        $this->ledger->balance -= $this->amount;
    }
}
