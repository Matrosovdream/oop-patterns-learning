<?php

declare(strict_types=1);

namespace App\Support;

use DateTimeImmutable;

/**
 * Clock abstracts "what time is it" so services don't call `new DateTimeImmutable`
 * directly. Inject a fixed clock in a test and time becomes deterministic.
 */
interface Clock
{
    public function now(): DateTimeImmutable;
}
