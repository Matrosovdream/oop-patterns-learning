<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use App\Notify\Broadcast;
use App\Notify\EmailSender;
use App\Notify\Sender;
use App\Notify\SmsSender;

$recipients = ['ada@example.com', 'linus@example.com'];

// Broadcast only knows the Sender abstraction. Hand it an EmailSender today…
$sender = new EmailSender(from: 'shop@shopkit.test');
echo "via email:\n";
Broadcast::to($sender, $recipients, 'Your order shipped!');

// …and swap in an SmsSender tomorrow with zero changes to Broadcast.
$sender = new SmsSender(gateway: 'twilio');
echo "via sms:\n";
Broadcast::to($sender, $recipients, 'Your order shipped!');

// The variable is typed to the interface, not a concrete class:
assert($sender instanceof Sender);
echo "\nSame Broadcast, two implementations — that's abstraction.\n";
