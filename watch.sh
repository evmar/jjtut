#!/bin/sh

set -e

./build.sh
fswatch -o book/*.go* book/*.css text | xargs -n1 ./build.sh
