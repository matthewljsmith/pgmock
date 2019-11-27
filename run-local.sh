#!/bin/bash -e

# grab script location
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

# re-build the app
$SCRIPT_DIR/build.sh

# run the local build
./pgmock.bin --verbose