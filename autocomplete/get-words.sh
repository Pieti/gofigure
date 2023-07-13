#!/usr/bin/env bash

URL="https://github.com/first20hours/google-10000-english/raw/master/google-10000-english.txt"
OUTPUT_FILE="vocabulary.txt"

wget "$URL" -O "$OUTPUT_FILE"

