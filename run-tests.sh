#!/bin/bash -e

# grab script location
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

pushd $SCRIPT_DIR/src
go test $1 ./...
popd