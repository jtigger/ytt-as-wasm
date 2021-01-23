#!/usr/bin/env bash

SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$SCRIPT_PATH/../build"

echo "Consider ./hack/watch.sh in another process for continuous build of WASM module."
echo ""
echo "open http://localhost:8080/"
go run ../web/server.go

