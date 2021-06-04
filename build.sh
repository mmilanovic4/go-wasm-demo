#!/bin/bash
#
# Build WASM

GO111MODULE=off GOOS=js GOARCH=wasm go build -o main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
