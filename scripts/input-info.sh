#!/bin/bash

echo "// auto-generated table, do not edit"
echo ".Liste des Fichiers de Tests"
echo "|==="
echo "|Fichier ^|Taille (octets) ^|Entropie (bits/octet)"
find . -type f -iname "*.decoded" | sort | while read FILE; do
    DATA=$(ent -t "$FILE" | tail -n +2 | tr ',' ' ')
    FILENAME=$(echo "$FILE" | sed 's/\.\///g')
    SIZE=$(echo "$DATA" | awk '{print $2}')
    ENTROPY=$(echo "$DATA" | awk '{print $3}')

    echo ""
    echo "| \`$FILENAME\`"
    echo "^| $SIZE"
    echo "^| $ENTROPY"
done
echo "|==="
