#!/usr/bin/env bash

mkdir ./dist
go build -o ./dist/sfs-macos
cp -r ./www ./dist/www