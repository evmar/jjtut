#!/bin/sh

set -e

fswatch -o book text | xargs -n1 ./build.sh
