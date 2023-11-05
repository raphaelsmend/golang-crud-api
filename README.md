<p align="center"><a href="https://go.dev/" target="_blank"><img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" width="400"></a></p>

# Register of Users

## Technical test for developing an full api users. this project was developed using containers. in the Golang language, using mysql as a database.

<!--ts-->
* [Frameworks used](#frameworks_used)
* [Prerequisites](#prerequisites)
* [Cloning repository](#clonning)
* [Installing](#installing)
* [Running](#running)
* [Swagger](#swagger)
<!--te-->

Frameworks used
============
- <a hfef="https://gin-gonic.com" target="_blank">Gin</a>
- <a hfef="https://github.com/gin-contrib/cors" target="_blank">Gin Cors</a>
- <a hfef="https://gorm.io/index.html" target="_blank">Gorm</a>
- <a hfef="https://pkg.go.dev/gopkg.in/validator.v2" target="_blank">Validator v2</a>
- <a hfef="https://github.com/dgrijalva/jwt-go" target="_blank">JWT</a>
- <a hfef="https://pkg.go.dev/github.com/joho/godotenv" target="_blank">Godotenv</a>

Prerequisites
============

```bash
- Linux (or WSL on windows)
- latest version of docker and docker-compose
```

Cloning repository
============

```bash
$ git clone git@github.com:raphaelsmend/golang-crud-api.git
```
Installing
============

```bash
$ golang-crud-api
$ docker compose build
```
Running
============
for start all containers
```bash
$ docker compose up

* this command runs the mysql server and start go server.
```

Routes Documentetion
============
all routes was made available in Imnsonia file with routes. filename: Insomnia_golang_crud_api.json