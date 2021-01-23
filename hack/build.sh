#!/usr/bin/env bash

# cd (project root)
SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "$SCRIPT_PATH/.."

# clean
if [[ -e build ]]; then
  rm -rf build
fi
mkdir build

cp -R web/ build/
GOOS=js GOARCH=wasm go build -o build/main.wasm
# https://github.com/golang/go/wiki/WebAssembly#getting-started
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" build

