@echo off

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0;%OLDGOPATH%

go install server
