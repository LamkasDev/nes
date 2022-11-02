:: Settings
@echo off

:: Build
cd ..\cmd
go build -o ..\bin\nes.exe .
cd ..

:: Run
.\bin\nes.exe -d