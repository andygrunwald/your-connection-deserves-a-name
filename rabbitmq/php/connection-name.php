<?php
require_once __DIR__ . '/vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;

define('RABBITMQ_HOST', '127.0.0.1');
define('RABBITMQ_PORT', 5672);
define('RABBITMQ_USERNAME', 'guest');
define('RABBITMQ_PASSWORD', 'guest');
define('RABBITMQ_ADDRESS', 'amqp://' . RABBITMQ_USERNAME . ':' . RABBITMQ_PASSWORD . '@' . RABBITMQ_HOST . ':' . RABBITMQ_PORT . '/');
define('CONNECTION_NAME', 'stock-exchange-rates-app');
define('RABBITMQ_MANAGEMENT_UI', 'http://' . RABBITMQ_HOST . ':15672/#/connections');

// Doing all the magic: Setting a proper connection name.
//
// In RabbitMQ this is supported via the custom properties you can set via the AMQP config.
// This is part of the AMQP specification.
//
// The RabbitMQ Management UI displays the connection name in the Connection listing.
//
// RabbitMQ docs:
//	- Connections: Client-Provided Connection Name: https://www.rabbitmq.com/connections.html#client-provided-names
//
// Attention:
//      This seems to be a "hacky" solution by the library "php-amqplib/php-amqplib".
//      It is possible to change in future.
//      Please keep track of https://github.com/php-amqplib/php-amqplib/issues/728
AMQPStreamConnection::$LIBRARY_PROPERTIES['connection_name'] = array('S', CONNECTION_NAME);

// Building the connection
output('Connecting to RabbitMQ on ' . RABBITMQ_ADDRESS);
$connection = new AMQPStreamConnection(RABBITMQ_HOST, RABBITMQ_PORT, RABBITMQ_USERNAME,RABBITMQ_PASSWORD);
output('Connecting to RabbitMQ on ' . RABBITMQ_ADDRESS . '  ... Successful');

output('');
output('Keeping the connection open ...');
output('You can connect to the RabbitMQ management UI: ' . RABBITMQ_MANAGEMENT_UI . ' (Username: ' . RABBITMQ_USERNAME . ', Password: ' . RABBITMQ_PASSWORD . ')');
output('	-> Check the Connections Tab and see the Connection Name');
output('Hit CTRL + C or cancel the process to stop.');
output('');

// Just looping to keep the process running.
// In detail we do not keep the connection active, but rely on the standard timeout.
// This way we provide time for the user to check the connection name behaviour.
while(true) {
    sleep(5);
}

/**
 * outputs logs $message with date and time to stdout.
 *
 * @param string $message Message to write out to stdout
 * @return void
 */
function output($message) {
    echo date('Y/m/d H:i:s') . ' ' . $message . PHP_EOL;
}