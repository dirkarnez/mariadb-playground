@echo off
REM --no-defaults
REM %~dp0

cd mariadb-10.6.5-winx64

SET F=data
IF EXIST %F% RMDIR /S /Q %F%

cd bin
mariadb-install-db.exe && mariadbd.exe --console
pause
