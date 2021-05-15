#!/usr/bin/env bash

set -e

export GOOS=linux GOARCH=amd64

D=dist/linux

go build -o "$D/client" cmd/client.go

# scp "$D/client" server:/usr/local/bin/remote-run
# cp "$D/client" /usr/local/bin/remote-run
# podman cp /data/DevelopmentRoot/github.com/gongt/remote-shell/dist/linux/client qbittorrent:/bin/x-www-browser
