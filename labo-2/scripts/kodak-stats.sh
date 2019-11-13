#!/bin/bash

function get_psnr() {
    local original="$1"
    local modified="$2"

    local original_ppm
    local modified_ppm
    local psnr
    local psnr_y
    local psnr_cb
    local psnr_cr

    original_ppm=$(mktemp)
    modified_ppm=$(mktemp)
    rm "$original_ppm" "$modified_ppm"
    original_ppm="$original_ppm.ppm"
    modified_ppm="$modified_ppm.ppm"

    convert "$original" "$original_ppm"
    convert "$modified" "$modified_ppm"

    psnr=$(pnmpsnr "$original_ppm" "$modified_ppm" 2>&1 | sed 1d)
    psnr_y=$(echo "$psnr" | grep "Y:" | awk '{print $3}')
    psnr_cb=$(echo "$psnr" | grep "CB:" | awk '{print $3}')
    psnr_cr=$(echo "$psnr" | grep "CR:" | awk '{print $3}')
    rm "$original_ppm" "$modified_ppm"

    printf "%s %s %s" "$psnr_y" "$psnr_cb" "$psnr_cr"
}

function get_ssim() {
    local original="$1"
    local modified="$2"

    ssim=$(./scripts/ssim.py -f "$original" -s "$modified" 2> /dev/null)

    printf "%1.3f" "$ssim"
}

function get_compression() {
    local original="$1"
    local modified="$2"

    local size_original
    local size_modified
    local compression

    size_original=$(stat -c %s "$original")
    size_modified=$(stat -c %s "$modified")
    compression=$(echo "100*(1-($size_modified/$size_original))" | bc -l)
    compression=$(printf "%2.2f" "$compression")

    printf "%s" "$compression"
}

function output_header() {
    echo "// auto-generated table, do not edit"
    echo ".$TABLE_TITLE"
    echo "[%autowidth,width=100%]"
    echo "|==="
    echo ".3+^.^h|Fichier 5+^h|JPEG 5+^h|JPEG2000"
    echo ".2+^.^h|Compression 3+^h|PSNR .2+^.^h|SSIM .2+^.^h|Compression 3+^h|PSNR .2+^.^h|SSIM"
    echo "^h|Y ^h|Cb ^h|Cr ^h|Y ^h|Cb ^h|Cr"
}

function output_table() {
    local original_directory="$1"
    local jpeg_directory="$2"
    local jpeg2000_directory="$3"
    local jpeg2000_data_directory="$4"
    local test_names="${*:5}"

    local original_file
    local jpeg_file
    local jpeg2000_file

    local jpeg_psnr
    local jpeg2000_psnr

    local jpeg_compression
    local jpeg2000_compression

    for test_name in ${test_names[*]}; do
        if echo "$test_name" | grep scripts > /dev/null; then
            continue
        fi
        test_name=$(basename "$test_name")
        test_name="${test_name%.*}"

        original_file="$original_directory/$test_name.png"
        jpeg_file="$jpeg_directory/$test_name.jpg"
        jpeg2000_file="$jpeg2000_directory/$test_name.png"
        jpeg2000_data_file="$jpeg2000_data_directory/$test_name.data"

        jpeg_psnr=($(get_psnr "$original_file" "$jpeg_file"))
        jpeg2000_psnr=($(get_psnr "$original_file" "$jpeg2000_file"))

        jpeg_ssim=$(get_ssim "$original_file" "$jpeg_file")
        jpeg2000_ssim=$(get_ssim "$original_file" "$jpeg2000_file")

        jpeg_compression=$(get_compression "$original_file" "$jpeg_file")
        jpeg2000_compression=$(get_compression "$original_file" "$jpeg2000_data_file")

        echo ""
        echo ".^| \`$(basename "$test_name")\`"
        echo "^.^| $jpeg_compression %"
        echo "^.^| ${jpeg_psnr[0]}"
        echo "^.^| ${jpeg_psnr[1]}"
        echo "^.^| ${jpeg_psnr[2]}"
        echo "^.^| $jpeg_ssim"
        echo "^.^| $jpeg2000_compression %"
        echo "^.^| ${jpeg2000_psnr[0]}"
        echo "^.^| ${jpeg2000_psnr[1]}"
        echo "^.^| ${jpeg2000_psnr[2]}"
        echo "^.^| $jpeg2000_ssim"
    done
}

function output_footer() {
    echo "|==="
}

if [[ -z ${ORIGINAL_DIR+x} ]]; then
    echo "ORIGINAL_DIR variable is not set"
    exit 1
fi

if [[ -z ${JPEG_DIR+x} ]]; then
    echo "JPEG_DIR variable is not set"
    exit 1
fi

if [[ -z ${JPEG2000_DIR+x} ]]; then
    echo "JPEG2000_DIR variable is not set"
    exit 1
fi

if [[ -z ${JPEG2000_DATA_DIR+x} ]]; then
    echo "JPEG2000_DATA_DIR variable is not set"
    exit 1
fi

output_header
output_table "$ORIGINAL_DIR" "$JPEG_DIR" "$JPEG2000_DIR" "$JPEG2000_DATA_DIR" "${*}"
output_footer
