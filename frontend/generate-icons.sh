#!/bin/bash

# Create icons directory if it doesn't exist
mkdir -p public/icons

SIZES=(72 96 128 144 152 192 384 512)

for size in "${SIZES[@]}"; do
    convert -size ${size}x${size} xc:black \
    -fill white \
    -gravity center \
    -pointsize $(($size/3)) \
    -font DejaVu-Sans \
    -draw "text 0,0 'arTstK'" \
    "public/icons/icon-${size}x${size}.png"
done 