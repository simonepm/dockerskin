<?php

require __DIR__ . '/vendor/autoload.php';

use Simonepm\Argumentor\Command;
use Simonepm\Argumentor\Argument;

$command = new Command();

$command->RegisterArgument("name");

$command->Exec(function (Argument $argument = null) {
    echo "Hello " . $argument->Get("name") . PHP_EOL;
});
