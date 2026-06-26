#!/bin/sh

set -e

fswatch -o book/*.go* book/*.css text | xargs -n1 ./build.sh
