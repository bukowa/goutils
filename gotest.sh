#!/usr/bin/env bash
set -e
echo "Running go test in $1"
cd "$1" && go test -v ./...
