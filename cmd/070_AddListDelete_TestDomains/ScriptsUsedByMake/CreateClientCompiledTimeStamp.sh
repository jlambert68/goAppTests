#!/bin/bash

timestamp=`date --rfc-3339=seconds`

echo "package main" > cmd/070_AddListDelete_TestDomains/wasm/clientCompiled.go
echo "" >> cmd/070_AddListDelete_TestDomains/wasm/clientCompiled.go
echo "const clientCompiledTimeStamp = " \"$timestamp\" >> cmd/070_AddListDelete_TestDomains/wasm/clientCompiled.go

