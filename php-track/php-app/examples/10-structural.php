<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use App\Cart\Cart;
use App\Cart\Line;
use App\Money\Money;
use App\Notify\EmailSender;
use App\Pricing\PercentageOff;
use App\Structural\CachingRates;
use App\Structural\Checkout;
use App\Structural\LegacySms;
use App\Structural\PrefixSender;
use App\Structural\SenderGroup;
use App\Structural\SlowRates;
use App\Structural\SmsAdapter;

// Adapter: a legacy SMS client now behaves like a Sender.
$sms = new SmsAdapter(new LegacySms());

// Decorator: wrap it to tag messages — same interface, extra behavior.
$tagged = new PrefixSender($sms, '[ShopKit] ');

// Composite: email + (decorated) SMS treated as one Sender.
$group = new SenderGroup(new EmailSender('shop@shopkit.test'), $tagged);
echo "adapter + decorator + composite:\n";
$group->send('buyer@example.com', 'Your order shipped!');

// Proxy: the caching proxy hits the slow backend only once per currency.
$backend = new SlowRates();
$rates = new CachingRates($backend);
for ($i = 0; $i < 3; $i++) {
    $rates->rate('USD');
}
printf("\nproxy: 3 lookups, backend called %d time(s)\n", $backend->calls);

// Facade: one call hides the cart → discount → notify orchestration.
$cart = new Cart('USD');
$cart->add(new Line('MUG-1', 'Enamel Mug', Money::fromMajor(12, 0, 'USD'), 2));
$checkout = new Checkout(new EmailSender('orders@shopkit.test'));
$total = $checkout->placeOrder($cart, new PercentageOff(10), 'buyer@example.com');
printf("\nfacade: charged %s\n", $total);
