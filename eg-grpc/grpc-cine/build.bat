@echo off

set DIR=%~dp0
set PROTODIR=%~dp0proto


echo ---------------- create proto java files ---------------

cd /d %PROTODIR%

for /f "delims=" %%i in ('dir /b *.proto') do (
	echo ---------------- %%i
	protoc -I . --go_out=plugins=grpc:../ ./%%i
)

echo ---------------- finish build go proto file ---------------

pause
