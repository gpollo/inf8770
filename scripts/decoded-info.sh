#!/bin/bash

function print_row() {
    local test=${1%.*}
    local size_original
    local size_arithmetic
    local size_dictionnary

    size_original=$(stat --printf="%s" "$test.decoded")
    size_arithmetic=$(stat --printf="%s" "$test.encoded.arithmetic")
    size_dictionnary=$(stat --printf="%s" "$test.encoded.dictionnary")

    echo ""
    echo "| \`$test\`"
    echo "^| $size_original"
    if [[ "$size_arithmetic" -le "$size_dictionnary" ]]; then
        echo "^| *${size_arithmetic}*"
        echo "^| ${size_dictionnary}"
    else
        echo "^| ${size_arithmetic}"
        echo "^| *${size_dictionnary}*"
    fi
}

echo "// auto-generated table, do not edit"
echo ".Résultats des Fichiers de Tests"
echo "|==="
echo "|Fichier ^|Originale (octets) ^|Arithmétique (octets) ^|Dictionnaire (octets) "
find . -type f -iname "*.decoded" | sort | while read FILE; do
    print_row "$FILE"
done
echo "|==="
