#!/usr/bin/env bash

go build -o x-www-browser cmd/client.go
scp x-www-browser server:/var/lib/machines/qbittorrent/usr/local/bin
