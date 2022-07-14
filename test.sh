#!/bin/bash

echo "====> Test suite started"
go test -v -coverpkg=./... -coverprofile=coverage.out ./...                                                       
echo "====> Test suite finished"

echo "====> Applying coverage ignore patterns"
while read ignore_pattern
do
	sed -i "\|${ignore_pattern}|d" coverage.out
done < coverage.ignore

echo "====> Generating coverage report"
go tool cover -func coverage.out
