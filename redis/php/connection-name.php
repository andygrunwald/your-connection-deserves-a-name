<?php
require_once './vendor/autoload.php';

define('REDIS_ADDRESS', 'tcp://127.0.0.1:6379');
define('REDIS_CONNECTION_NAME', 'your-connection-deserves-a-name-php');

// Building the connection
output('Connecting to Redis on ' . REDIS_ADDRESS);
$client = new Predis\Client(REDIS_ADDRESS);
output('Connecting to Redis on ' . REDIS_ADDRESS . '  ... Successful');

// Doing all the magic: Setting a proper connection name.
//
// In redis this is supported via the command "CLIENT SETNAME <connection-name>".
// The name of the _current_ connection can be retrieved via "CLIENT GETNAME".
// To see all names connected to a redis instance, execute "CLIENT LIST".
// The connection name value is visible in the `name` field.
//
// Redis docs:
//	- CLIENT SETNAME: https://redis.io/commands/client-setname
//	- CLIENT GETNAME: https://redis.io/commands/client-getname
//	- CLIENT LIST: https://redis.io/commands/client-list
output('Setting client name "' . REDIS_CONNECTION_NAME . '"');
$client->client('SETNAME', REDIS_CONNECTION_NAME);
output('Setting client name "' . REDIS_CONNECTION_NAME . '" ... Successful');

output('');
output('Playing PING/PONG to keep connection open ...');
output('You can connect to redis and execute "CLIENT LIST" to see the connection');
output('Hit CTRL + C or cancel the process to stop.');
output('');

// Sending a ping to keep the connection open and provide time for
// the user to check the CLIENT SETNAME behaviour.
while(true) {
    output('PING ...');
    $val = $client->ping();
    output('     ... ' . $val->getPayload());

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