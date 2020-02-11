#! /bin/bash

if [ ! -d bin ]; then
    mkdir bin
fi

for cmd in cmd/*.go; do
    if [[ $cmd =~ cmd/(.+).go ]]; then
        go build -o bin/${BASH_REMATCH[1]} $cmd
    fi
done
