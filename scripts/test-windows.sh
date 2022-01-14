#!/usr/bin/env bash

set -Eeuo pipefail

export GOOS=windows GOARCH=amd64

D="dist/windows"

ExeName='remote-shell.server.test'

export GOTMPDIR="$(dirname "$(realpath "${BASH_SOURCE[0]}")")/../dist/temp"
go build  -o "$D/$ExeName.exe" "cmd/server-tray.go"

ls -lhA --color=auto "$D/$ExeName.exe"
