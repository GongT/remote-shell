#!/usr/bin/env pwsh
Set-StrictMode -Version latest
$ErrorActionPreference = "Stop"

$D = "dist/windows"

$ExeName = 'remote-shell.server.release'

go build -ldflags -H=windowsgui -o "$D/$ExeName.exe" "cmd\server-tray.go" || { exit 1 }

$T = "$env:APPDATA\Microsoft\Windows\Start Menu\Programs\Startup"

Stop-Process -Name $ExeName -ErrorAction SilentlyContinue

Write-Output "install to $T"
Copy-Item "$D/$ExeName.exe" $T
Start-Process "$T/$ExeName.exe"
