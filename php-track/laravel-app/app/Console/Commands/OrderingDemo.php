<?php

declare(strict_types=1);

namespace App\Console\Commands;

use App\Ordering\Application\PayOrderHandler;
use App\Ordering\Application\PlaceOrderCommand;
use App\Ordering\Application\PlaceOrderHandler;
use App\Ordering\Application\PlaceOrderLine;
use App\Ordering\Application\ShipOrderHandler;
use App\Ordering\Application\SummaryProjection;
use App\Ordering\Domain\OrderException;
use App\Ordering\Infrastructure\RecordingPublisher;
use Illuminate\Console\Command;

/**
 * `php artisan ordering:demo` — runs the whole bounded context end to end. The
 * handler arguments are auto-injected by the container (constructor/method
 * injection), which resolves the ports to the adapters bound in
 * OrderingServiceProvider. This command is the presentation layer for the CLI.
 */
final class OrderingDemo extends Command
{
    protected $signature = 'ordering:demo';

    protected $description = 'Run the Ordering bounded context end to end (lessons 12-20).';

    public function handle(
        PlaceOrderHandler $place,
        PayOrderHandler $pay,
        ShipOrderHandler $ship,
        SummaryProjection $readModel,
        RecordingPublisher $events,
    ): int {
        // Happy path: place -> pay -> ship.
        $this->info('== ORD-1: full lifecycle ==');
        $id = $place->handle(new PlaceOrderCommand('ORD-1', 'USD', [
            new PlaceOrderLine('MUG-1', 'Enamel Mug', 1200, 2),
            new PlaceOrderLine('TEE-1', 'Logo Tee', 2500, 1),
        ]));
        $pay->handle((string) $id);
        $ship->handle((string) $id);

        // Second order left pending.
        $this->info('== ORD-2: placed, left pending ==');
        $place->handle(new PlaceOrderCommand('ORD-2', 'USD', [
            new PlaceOrderLine('CAP-1', 'Beanie', 1800, 1),
        ]));

        // Invalid: an empty order is rejected by a domain invariant.
        $this->info('== ORD-3: invalid (no lines) ==');
        try {
            $place->handle(new PlaceOrderCommand('ORD-3', 'USD', []));
        } catch (OrderException $e) {
            $this->line('  correctly rejected: ' . $e->getMessage());
        }

        $this->newLine();
        $this->info('domain events (captured by the RecordingPublisher adapter):');
        foreach ($events->names() as $name) {
            $this->line('  - ' . $name);
        }

        $this->newLine();
        $this->info('order book (CQRS read model, projected from those events):');
        $this->table(
            ['id', 'status', 'total'],
            array_map(static fn ($summary) => $summary->toArray(), $readModel->all()),
        );

        return self::SUCCESS;
    }
}
