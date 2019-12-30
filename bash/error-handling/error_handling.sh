#!/usr/bin/env bash

function error_exit {
    echo "Usage: ./error_handling <greetee>" 1>&2
    exit 1
}

if (( $# != 1 ))
then
    error_exit
fi

echo "Hello, $1"