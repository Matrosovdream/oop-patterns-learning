<?php

declare(strict_types=1);

namespace App\Structural;

use App\Notify\Sender;

/**
 * Composite — a group of Senders that is itself a Sender. send() fans out to
 * every child, so callers treat one or many identically.
 */
final class SenderGroup implements Sender
{
    /** @var list<Sender> */
    private array $children;

    public function __construct(Sender ...$children)
    {
        $this->children = $children;
    }

    public function send(string $to, string $message): void
    {
        foreach ($this->children as $child) {
            $child->send($to, $message);
        }
    }
}
