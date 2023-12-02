#!/bin/zsh

year=$1
day=$2

dir_path="$year/day$(printf %02d $day)"

curl -s -b $(cat .session_cookie) \
    "https://adventofcode.com/$year/day/$day/input" \
    -o "$dir_path/input.txt"
