<?php

declare(strict_types=1);

namespace App\Structural;

/** Looks up a tax rate as a percentage. Pretend it's a slow remote call. */
interface RateProvider
{
    public function rate(string $currency): int;
}
