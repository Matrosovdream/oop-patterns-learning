<?php

declare(strict_types=1);

namespace App\Http\Controllers;

use App\Ordering\Application\PayOrderHandler;
use App\Ordering\Application\PlaceOrderCommand;
use App\Ordering\Application\PlaceOrderHandler;
use App\Ordering\Application\PlaceOrderLine;
use App\Ordering\Application\ShipOrderHandler;
use App\Ordering\Application\SummaryProjection;
use Illuminate\Http\JsonResponse;

/**
 * Presentation layer for HTTP. It does no business logic — it builds a command,
 * calls use cases (auto-injected by the container), and shapes a response. The
 * exact same use cases back the CLI command (OrderingDemo); only the presentation
 * differs. That reuse is the whole point of a thin application layer.
 */
final class OrderingController extends Controller
{
    public function demo(
        PlaceOrderHandler $place,
        PayOrderHandler $pay,
        ShipOrderHandler $ship,
        SummaryProjection $readModel,
    ): JsonResponse {
        $id = $place->handle(new PlaceOrderCommand('WEB-1', 'USD', [
            new PlaceOrderLine('MUG-1', 'Enamel Mug', 1200, 2),
            new PlaceOrderLine('TEE-1', 'Logo Tee', 2500, 1),
        ]));
        $pay->handle((string) $id);
        $ship->handle((string) $id);

        return response()->json([
            'message' => 'Placed, paid and shipped WEB-1. Below is the CQRS read model.',
            'orders' => array_map(static fn ($summary) => $summary->toArray(), $readModel->all()),
        ]);
    }
}
