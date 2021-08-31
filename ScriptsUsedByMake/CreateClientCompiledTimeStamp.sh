#!/bin/bash

timestamp=`date --rfc-3339=seconds`

echo "package main" > cmd/060_golangWasmMultiProto_goAppv8/wasm/clientCompiled.go
echo "" >> clientCompile.go
echo "const clientCompiledTimeStamp = " \"$timestamp\" >> clientCompile.go