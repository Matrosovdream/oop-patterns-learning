<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use App\Catalog\CatalogService;
use App\Catalog\InMemoryProductRepository;
use App\Catalog\ProductRepository;
use App\Money\Money;
use App\Notify\EmailSender;
use App\Notify\Sender;
use App\Support\Clock;
use App\Support\SystemClock;

// ---- Composition root ----
// This script is the only place that picks concrete implementations. It builds
// the low-level details and injects them into the high-level CatalogService.
// Swap any line below and CatalogService is unaffected.
$repository = new InMemoryProductRepository(); // a ProductRepository
$notifier = new EmailSender('catalog@shopkit.test'); // a Sender
$clock = new SystemClock(); // a Clock

// Type-checks against the interfaces, proving the service depends only on abstractions:
assert($repository instanceof ProductRepository);
assert($notifier instanceof Sender);
assert($clock instanceof Clock);

$service = new CatalogService($repository, $notifier, $clock);

// ---- Use the wired application ----
$service->addProduct('MUG-1', 'Enamel Mug', Money::fromMajor(12, 0, 'USD'));
$service->addProduct('TEE-1', 'Logo Tee', Money::fromMajor(25, 0, 'USD'));

printf("\ncatalog now has %d products.\n", $service->count());
echo "CatalogService never named a database or a mailer — only abstractions.\n";
