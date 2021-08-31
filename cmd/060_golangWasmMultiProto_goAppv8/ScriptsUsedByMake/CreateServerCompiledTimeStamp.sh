#!/bin/bash

timestamp=`date --rfc-3339=seconds`

echo "package main" > cmd/060_golangWasmMultiProto_goAppv8/webServer/serverCompiled.go
echo "" >> cmd/060_golangWasmMultiProto_goAppv8/webServer/serverCompiled.go
echo "const serverCompiledTimeStamp = " \"$timestamp\" >> cmd/060_golangWasmMultiProto_goAppv8/webServer/serverCompiled.go
