#!/bin/bash

timestamp=`date --rfc-3339=seconds`

echo "package main" > cmd/060_golangWasmMultiProto_goAppv8/wasm/clientCompiled.go
echo "" >> cmd/060_golangWasmMultiProto_goAppv8/wasm/clientCompiled.go
echo "const clientCompiledTimeStamp = " \"$timestamp\" >> cmd/060_golangWasmMultiProto_goAppv8/wasm/clientCompiled.go

