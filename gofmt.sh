#!/usr/bin/env bash
set -e
echo "Running go fmt in $1"
cd "$1" && go fmt ./...
