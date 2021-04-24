#!/usr/bin/env bash

set -e

export GOOS=linux GOARCH=arm GOARM=7

D=dist/raspi

go build -o $D/client cmd/client.go
