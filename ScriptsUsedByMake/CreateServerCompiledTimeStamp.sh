#!/bin/bash

timestamp=`date --rfc-3339=seconds`

echo $timestamp

echo "package main" > clientCompile.go
echo "" >> clientCompile.go
echo "const clientCompiledTimeStamp = " \"$timestamp\" >> clientCompile.go