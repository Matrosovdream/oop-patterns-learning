<?php

declare(strict_types=1);

namespace App\Behavioral;

/** The receiver the commands act on. */
final class Ledger
{
    public function __construct(public int $balance = 0)
    {
    }
}
