#!/bin/bash

######################
# colorspace columns #
######################

function output_header_colorspace() {
    echo "^|Espace de Couleurs"
}


function output_table_colorspace() {
    local filename=$1
    local colorspace

    colorspace=$(basename "$filename")
    colorspace=${colorspace%.*}
    colorspace=${colorspace^^}
    
    echo "^| $colorspace"
}

##########################
# haar_recursion columns #
##########################

function output_header_haar_recursion() {
    echo "^|Récursion"
}


function output_table_haar_recursion() {
    local filename=$1
    local recursion

    recursion=$(basename "$filename")
    recursion=${recursion%.*}
    recursion=${recursion//haar-/}
    
    echo "^| $recursion"
}

##########################
# deadzone_delta columns #
##########################

function output_header_deadzone_delta() {
    echo "^|Pas"
}


function output_table_deadzone_delta() {
    local filename=$1
    local delta

    delta=$(basename "$filename")
    delta=${delta%.*}
    delta=${delta//deadzone-delta-/}
    
    echo "^| $delta"
}

##########################
# deadzone_width columns #
##########################

function output_header_deadzone_width() {
    echo "^|Zone"
}


function output_table_deadzone_width() {
    local filename=$1
    local width

    width=$(basename "$filename")
    width=${width%.*}
    width=${width//deadzone-width-/}
    
    echo "^| $width"
}

############################
# deadzone_various columns #
############################

function output_header_deadzone_various() {
    echo "^|Zone ^|Pas"
}


function output_table_deadzone_various() {
    local filename=$1
    local width
    local delta

    filename=$(basename "$filename")
    filename=${filename%.*}
    width=${filename//deadzone-various-/}
    width=${width//-*/}
    delta=${filename//deadzone-various-/}
    delta=${delta//*-/}
    
    echo "^| $width"
    echo "^| $delta"
}

#################
# mixed columns #
#################

function output_header_mixed() {
    echo "^|Ondelette"
}


function output_table_mixed() {
    local filename=$1
    local wavelet

    filename=$(basename "$filename")
    filename=${filename%.*}
    filename=${filename//mixed-/}
    
    if echo "$filename" | grep "haar" > /dev/null; then
        wavelet=${filename//haar/}
        wavelet="Haar de récursion $wavelet"
    fi

    if echo "$filename" | grep "daub" > /dev/null; then
        wavelet="Daubechies"
    fi

    echo "^| $wavelet"
}

###########################
# midthread_delta columns #
###########################

function output_header_midthread_delta() {
    echo "^|Taille du Pas"
}


function output_table_midthread_delta() {
    local filename=$1
    local delta

    delta=$(basename "$filename")
    delta=${delta%.*}
    delta=${delta//midthread-delta-/}
    
    echo "^| $delta"
}

##################
# normal columns #
##################

function output_header() {
    local additional_header=""
    local additional_header_function

    if [[ -n ${ADDITIONAL_COLUMNS+x} ]]; then
        additional_header_function="output_header_$ADDITIONAL_COLUMNS"
        additional_header=$($additional_header_function)
        additional_header="$additional_header "
    fi

    echo "// auto-generated table, do not edit"
    echo ".$TABLE_TITLE"
    echo "[%autowidth,width=100%]"
    echo "|==="
    echo "|Fichier $additional_header^|Taille (octets) ^|Compression"
}


function output_table() {
    local original="$1"
    local files=${*:2}

    local size_original
    local size_encoded
    local compression

    local additional_output_function

    size_original=$(stat -c %s "$original")
    for encoded in ${files[*]}; do
        if echo "$encoded" | grep scripts > /dev/null; then
            continue
        fi

        size_encoded=$(stat -c %s "$encoded")
        compression=$(echo "100*(1-($size_encoded/$size_original))" | bc -l)
        compression=$(printf "%2.2f" "$compression")

        echo ""
        echo "| \`$(basename "$encoded")\`"

        if [[ -n ${ADDITIONAL_COLUMNS+x} ]]; then
            additional_output_function="output_table_$ADDITIONAL_COLUMNS"
            "$additional_output_function" "$encoded"
        fi

        echo "^| $size_encoded"
        echo "^| $compression %"
    done
}

function output_footer() {
    echo "|==="
}

if [[ -z ${ORIGINAL_IMAGE+x} ]]; then
    echo "ORIGINAL_IMAGE variable is not set"
    exit 1
fi

output_header
output_table "$ORIGINAL_IMAGE" "${*}"
output_footer
