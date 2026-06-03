<?php

declare(strict_types=1);

namespace App\Providers;

use App\Ordering\Application\EventPublisher;
use App\Ordering\Application\SummaryProjection;
use App\Ordering\Domain\OrderRepository;
use App\Ordering\Infrastructure\FanOutPublisher;
use App\Ordering\Infrastructure\InMemoryOrderRepository;
use App\Ordering\Infrastructure\LoggingPublisher;
use App\Ordering\Infrastructure\RecordingPublisher;
use Illuminate\Contracts\Foundation\Application;
use Illuminate\Support\ServiceProvider;

/**
 * The COMPOSITION ROOT for the Ordering context. Laravel's service container is
 * our dependency-injection engine; this provider is the one place that maps the
 * inner layers' PORTS (interfaces) to concrete outer-layer ADAPTERS. Change a
 * binding here and the use cases — which only know the interfaces — don't change.
 * That's dependency inversion (lesson 08) realized at application scale (lessons
 * 13–14).
 */
final class OrderingServiceProvider extends ServiceProvider
{
    public function register(): void
    {
        // Repository port -> in-memory adapter. Singleton so state persists across
        // a single request/command (swap for an Eloquent adapter and nothing else
        // changes).
        $this->app->singleton(OrderRepository::class, InMemoryOrderRepository::class);

        // Read model + recorder as singletons so we can query them after handling.
        $this->app->singleton(SummaryProjection::class);
        $this->app->singleton(RecordingPublisher::class);

        // EventPublisher port -> a FanOut composite over log + recorder + read model,
        // so one event stream drives all three (lesson 19's CQRS projection).
        $this->app->singleton(EventPublisher::class, static function (Application $app): EventPublisher {
            return new FanOutPublisher(
                new LoggingPublisher(),
                $app->make(RecordingPublisher::class),
                $app->make(SummaryProjection::class),
            );
        });
    }
}
