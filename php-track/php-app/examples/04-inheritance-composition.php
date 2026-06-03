<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use App\Catalog\Product;
use App\Money\Money;
use App\Notify\EmailSender;
use App\Notify\Sender;

$now = new DateTimeImmutable();

// Product USES the Timestamps trait, so it has touch()/createdAt()/updatedAt()
// "for free" — horizontal reuse, no inheritance.
$mug = Product::create('MUG-1', 'Enamel Mug', Money::fromMajor(12, 0, 'USD'), $now);
printf("created %s at %s\n", $mug->name, $mug->createdAt()?->format('H:i'));

$mug->reprice(Money::fromMajor(9, 50, 'USD'), $now->modify('+1 hour'));
printf("repriced to %s; updated at %s\n", $mug->price(), $mug->updatedAt()?->format('H:i'));

// PriceWatcher COMPOSES a Sender by holding it as a field. It reuses sending
// behavior WITHOUT being a Sender — and swapping email for SMS won't change it.
// (Contrast: making PriceWatcher `extends EmailSender` would weld it to email
// forever and expose EmailSender's whole surface. Composition keeps it free.)
final class PriceWatcher
{
    public function __construct(private readonly Sender $sender)
    {
    }

    public function announce(Product $product, string $to): void
    {
        $this->sender->send($to, sprintf('%s is now %s', $product->name, $product->price()));
    }
}

$watcher = new PriceWatcher(new EmailSender('deals@shopkit.test'));
$watcher->announce($mug, 'fan@example.com');

echo "\nTrait = horizontal reuse; field = composition. Inheritance is reserved for true 'is-a'.\n";
