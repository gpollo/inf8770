#!/bin/bash

function test_speed_arithmetic() {
    local workers=$1
    local data
    local time

    data=$(mktemp)
    cat /dev/urandom | head -c 1000 > "$data"
    time=$(cat "$data" | ./bin/arithmetic --encode --benchmark --parallel --workers "$workers" | sed 's/us//g')

    printf "%d %s\n" "$workers" "$time"
}


test_speed_arithmetic 1
test_speed_arithmetic 4
test_speed_arithmetic 8
test_speed_arithmetic 16
test_speed_arithmetic 32
test_speed_arithmetic 64
