<?php

declare(strict_types=1);

namespace App\Behavioral;

/** A request packaged as an object — so it can be queued, logged, or undone. */
interface Command
{
    public function execute(): void;

    public function undo(): void;
}
