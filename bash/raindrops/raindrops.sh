#!/usr/bin/env bash

let pling=$1%3
let plang=$1%5
let plong=$1%7

response=""

if (( $pling == 0 ))
then
    response+="Pling"
fi
if (( $plang == 0 ))
then
    response+="Plang"
fi
if (( $plong == 0 ))
then
    response+="Plong"
fi
if [[ $response == "" ]]
then
    echo "$1"
else
    echo $response
fi