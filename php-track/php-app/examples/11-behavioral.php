<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use App\Behavioral\Deposit;
use App\Behavioral\Event;
use App\Behavioral\IllegalTransitionException;
use App\Behavioral\Invoker;
use App\Behavioral\Ledger;
use App\Behavioral\OrderState;
use App\Behavioral\PendingState;
use App\Behavioral\Publisher;
use App\Money\Money;

// Strategy as a closure (PHP's lightest form): pick an algorithm at runtime.
/** @var array<string, callable(int): Money> $shipping */
$shipping = [
    'flat' => fn (int $grams): Money => Money::fromMajor(5, 0, 'USD'),
    'weight' => fn (int $grams): Money => Money::fromMajor(2, 0, 'USD')->multiply(intdiv($grams + 999, 1000)),
];
echo "strategy:\n";
foreach (['flat', 'weight'] as $name) {
    printf("  %-7s 1500g -> %s\n", $name, $shipping[$name](1500));
}

// Observer: two subscribers react to one published event.
$publisher = new Publisher();
$publisher->subscribe(fn (Event $e) => printf("  [email] order %s: %s\n", $e->orderId, $e->name));
$publisher->subscribe(fn (Event $e) => printf("  [audit] order %s: %s\n", $e->orderId, $e->name));
echo "observer:\n";
$publisher->publish(new Event('placed', 'A-100'));

// Command: run with full undo support.
$ledger = new Ledger();
$invoker = new Invoker();
$invoker->run(new Deposit($ledger, 100));
$invoker->run(new Deposit($ledger, 50));
printf("command: balance after 2 deposits = %d\n", $ledger->balance);
$invoker->undoLast();
printf("command: balance after undo       = %d\n", $ledger->balance);

// Template Method: a fixed frame with a varying body (passed as a callable).
$renderReport = static fn (string $title, callable $body): string => "== {$title} ==\n" . $body() . "\n-- end --";
echo "template:\n";
echo $renderReport('Daily Sales', fn (): string => "  units: 42\n  revenue: \$1,337") . "\n";

// State: legal transitions only.
echo "state:\n";
$state = new PendingState();
$state = step($state, 'ship'); // illegal: unpaid
$state = step($state, 'pay');  // pending -> paid
$state = step($state, 'ship'); // paid -> shipped
$state = step($state, 'pay');  // illegal: already shipped

function step(OrderState $state, string $action): OrderState
{
    try {
        $next = $action === 'pay' ? $state->pay() : $state->ship();
        printf("  %-7s from %-8s -> %s\n", $action, $state->name(), $next->name());

        return $next;
    } catch (IllegalTransitionException $e) {
        printf("  %-7s from %-8s -> rejected (%s)\n", $action, $state->name(), $e->getMessage());

        return $state;
    }
}
