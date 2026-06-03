<?php

declare(strict_types=1);

namespace App\Ordering\Domain;

/**
 * OrderStatus is a backed enum — the idiomatic PHP way to model a closed set of
 * values. Better than string constants: the type system guarantees a status is
 * always one of these four.
 */
enum OrderStatus: string
{
    case Draft = 'draft';     // being assembled; lines can be added
    case Pending = 'pending'; // placed, awaiting payment
    case Paid = 'paid';       // paid, awaiting shipment
    case Shipped = 'shipped'; // done
}
