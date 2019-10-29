#!/bin/bash

function download_file() {
    local path="$1"
    local filename
    local url

    filename=$(basename "$path")
    local url="http://r0k.us/graphics/kodak/kodak/$filename"

    curl --silent "$url" -o "$path" > /dev/null
}

download_file "$1"
