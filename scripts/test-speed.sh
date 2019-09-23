#!/bin/bash

function test_speed_arithmetic() {
    local file_size=$1
    local data
    local time

    data=$(mktemp)
    cat /dev/urandom | head -c "$file_size" > "$data"
    time=$(cat "$data" | ./bin/arithmetic --encode --benchmark --parallel --workers 16 | sed 's/us//g')

    printf "%d %s\n" "$file_size" "$time"
}

function test_speed_dictionnary() {
    local file_size=$1
    local data
    local time

    data=$(mktemp)
    cat /dev/urandom | head -c "$file_size" > "$data"
    time=$(cat "$data" | ./bin/dictionnary --encode --benchmark | sed 's/us//g')

    printf "%d %s\n" "$file_size" "$time"
}

TEST_CASES=$(printf "5 10 50 100 500 1000 2500" | tr ' ' '\n')

echo "$TEST_CASES" | while read N; do
    test_speed_arithmetic $N
done


echo "$TEST_CASES" | while read N; do
    test_speed_dictionnary $N
done
