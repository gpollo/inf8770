#!/bin/bash

function output_header() {
    echo "// auto-generated table, do not edit"
    echo ".$TABLE_TITLE"
    echo "[%autowidth,cols=4,frame=none,grid=none,stripe=all]"
    echo "|==="
}

function output_table() {
    local files=${*}
    local filename
    local queue

    # TODO: that's not gonna work with image count different than a multiple of 4

    queue=()
    for encoded in ${files[*]}; do
        if echo "$encoded" | grep scripts > /dev/null; then
            continue
        fi

        filename=${encoded//various-/}
        filename=${filename//png/bmp}
        echo "^| \`$(basename "$filename")\`"

        queue+=("$encoded")
        if [[ "${#queue[@]}" == "4" ]]; then
            for image in ${queue[*]}; do
                echo "^|image:${image}[width=140]"
            done
            queue=()
        fi
    done

}

function output_footer() {
    echo "|==="
}

output_header
output_table "${*}"
output_footer
