<?php

declare(strict_types=1);

namespace App\Behavioral;

/** Runs commands and can undo the most recent one. */
final class Invoker
{
    /** @var list<Command> */
    private array $history = [];

    public function run(Command $command): void
    {
        $command->execute();
        $this->history[] = $command;
    }

    public function undoLast(): void
    {
        $last = array_pop($this->history);
        $last?->undo();
    }
}
