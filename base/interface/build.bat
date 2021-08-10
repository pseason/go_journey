@echo off

@REM build go program on linux

set GOARCH=amd64
set GOOS=linux

go build .\ig1.go

pause