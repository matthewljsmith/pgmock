#!/bin/bash -e

# grab script location
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

# push source, build and return
pushd src
go build -o ../pgmock.bin .
popd