#!/usr/bin/env bash
set -e

TEMPLATE="Makefile.template"
echo "Generating Makefile for submodule $1 from $TEMPLATE"
printf "$(cat $TEMPLATE)" > "$1/Makefile"
