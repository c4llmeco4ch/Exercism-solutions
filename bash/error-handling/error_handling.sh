#!/usr/bin/env bash

status=0
output=""

function error_exit {
    echo "Usage: ./error_handling <greetee>"
    exit $status
}

if (( $# != 1 ))
then
    status=1
    error_exit
fi

echo "Hello, $1"