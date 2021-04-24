#!/usr/bin/env bash

set -e

export GOOS=linux GOARCH=arm GOARM=7

D=dist/router

go build -o $D/x-www-browser cmd/x-www-browser.go
go build -o $D/remote-run cmd/remote-run.go
go build -o $D/handler cmd/handler.go

scp $D/* router:/usr/local/bin
