@echo off

cd %~dp0

docker-compose up --build && docker-compose down
pause

