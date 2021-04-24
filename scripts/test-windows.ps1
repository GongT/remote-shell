#!/usr/bin/env pwsh
Set-StrictMode -Version latest
$ErrorActionPreference = "Stop"

$D = "dist/windows"
$ExeName = 'remote-shell.server.test'

$ExeNameRel = 'remote-shell.server.release'
Stop-Process -Name $ExeNameRel -ErrorAction SilentlyContinue

$env:GOTMPDIR = "$PSScriptRoot/../dist/temp"
New-Item -Type Directory $env:GOTMPDIR -ErrorAction SilentlyContinue

go build -o "$D/$ExeName.exe" "cmd\server.go" || { exit 1 }

& "$D/$ExeName.exe"
