#!/bin/zsh

year=$1
day=$2

curl -s -b $(cat .session_cookie) "https://adventofcode.com/$year/day/$day/input" -o "$year/day$(printf %02d $day)/input.txt"
