#!/usr/bin/env bash

SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "$SCRIPT_PATH/.."

ls *.go | entr -r -d -c bash -c 'echo -n "$( date +'%H:%M:%S' ) — Rebuilding main.wasm..." && GOOS=js GOARCH=wasm go build -o build/main.wasm && echo -n "done; refresh browser."'
