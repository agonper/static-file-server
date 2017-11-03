#!/usr/bin/env bash

mkdir ./dist
go build -o ./dist/sfs
cp -r ./www ./dist/www