#!/bin/bash

timestamp=`date --rfc-3339=seconds`

echo "package main" > cmd/070_AddListDelete_TestDomains/webServer/serverCompiled.go
echo "" >> cmd/070_AddListDelete_TestDomains/webServer/serverCompiled.go
echo "const serverCompiledTimeStamp = " \"$timestamp\" >> cmd/070_AddListDelete_TestDomains/webServer/serverCompiled.go
