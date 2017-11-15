@echo off
title GoMine : A Minecraft Bedrock Edition server software in Go
set loop=false

SET GOPATH=%~dp0
set /A "loops=1"

where go >nul 2>&1 && goto startScript || (
echo You require Go / Golang to run this program!
echo Download it from https://golang.org/ and try again!
pause>nul & exit
)

:startScript
powershell go run ./src/main.go
if /i "%loop%"=="true" (
    set /A "loops=loops + 1"
    echo Restarted %loops% time^(s^)
    echo To escape the loop, press CTRL+C now. Otherwise, wait 5 seconds for the server to restart.
    echo.
    ping localhost -n 5 >nul
    goto startScript
)
pause>nul & exit