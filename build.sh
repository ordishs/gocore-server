#!/bin/sh

cd $(dirname $BASH_SOURCE)

rm -rf build

mkdir -p build
mkdir -p build/windows
mkdir -p build/linux
mkdir -p build/raspian

PROG_NAME=gocore_server

go build -o build/darwin/${PROG_NAME}
env GOOS=linux GOARCH=amd64 go build -o build/linux/${PROG_NAME}
env GOOS=linux GOARCH=arm go build -o build/raspian/${PROG_NAME}
env GOOS=windows GOARCH=386 go build -o build/windows/${PROG_NAME}



