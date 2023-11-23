mariadb-playground
==================
From https://github.com/peter-evans/docker-compose-actions-workflow

[![Actions Status](https://github.com/dirkarnez/github-docker-compose-action/workflows/docker-compose-actions-workflow/badge.svg)](https://github.com/dirkarnez/github-docker-compose-action/actions)

### Prerequisites
- Unzip tool
- MariaDB client
   - [DBeaver](https://dbeaver.io/files/dbeaver-ce-latest-win32.win32.x86_64.zip), or
   - [HeidiSQL x64 Portable](https://www.heidisql.com/download.php?download=portable-64), or
- [Download MariaDB Server - MariaDB.org](https://mariadb.org/download/?t=mariadb&p=mariadb&r=10.6.5&os=windows&cpu=x86_64&pkg=zip&m=xtom_hk)
- MYSQL
   - https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-8.0.33-winx64.zip

### How to use
1. Download / `git clone` this repo
2. Unzip
3. Run [`local-dev.cmd`](local-dev.cmd) or [`local-docker-dev.cmd`](local-docker-dev.cmd) for initialization
4. Run [`mariadb-playground.exe  --database=default --docker=true`](mariadb-playground.exe) or `./mariadb-playground --database=default --docker=true` in docker every time to (re)create `DATABASE` (database name "EIE3112")
5. Connection string "root:@tcp(localhost:3306)/?charset=utf8&parseTime=True"

### TODOs
- test cases
- database name as [`mariadb-playground.exe`](mariadb-playground.exe)'s argument
- Performance tuning (Query explain)
  - [MySQL :: MySQL Workbench Manual :: 7.1 Performance Dashboard](https://dev.mysql.com/doc/workbench/en/wb-performance-dashboard.html)
  - [DBeaver Documentation](https://dbeaver.com/docs/wiki/Query-Execution-Plan/)
   - <kbd>Ctrl</kbd> + <kbd>Shift</kbd> + <kbd>E</kbd> to explain current query
- Indexing
- Faking
  - [pioz/faker: Random fake data and struct generator for Go.](https://github.com/pioz/faker)
- SQL Injection
   - [Hacking Websites with SQL Injection - Computerphile - YouTube](https://www.youtube.com/watch?v=_jKylhJtPmI)
- GORM
   - Join without raw string (`Preloads`)
      - [预加载 | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/docs/preload.html)
   	- db.Joins("Company", DB.Where(&Company{Alive: true})).Find(&users)
   - [gorm package - gorm.io/gorm - Go Packages](https://pkg.go.dev/gorm.io/gorm#section-readme)
