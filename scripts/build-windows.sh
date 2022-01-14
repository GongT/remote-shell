#!/usr/bin/env bash

set -Eeuo pipefail

export GOOS=windows GOARCH=amd64

D="dist/windows"

ExeName='remote-shell.server.release'

go build -ldflags -H=windowsgui -o "$D/$ExeName.exe" "cmd/server-tray.go"

ls -lhA --color=auto "$D/$ExeName.exe"
