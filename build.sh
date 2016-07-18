#!/bin/sh

mkdir -p ./bin
go get ./src
go build -o bin/wand ./src
