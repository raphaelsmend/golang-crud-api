<p align="center"><a href="https://laravel.com" target="_blank"><img src="https://raw.githubusercontent.com/laravel/art/master/logo-lockup/5%20SVG/2%20CMYK/1%20Full%20Color/laravel-logolockup-cmyk-red.svg" width="400"></a></p>

# Register of Users

## Technical test for developing an full api users. this project was developed using containers. in the Golang language, using mysql as a database.

<!--ts-->
* [Prerequisites](#prerequisites)
* [Cloning repository](#clonning)
* [Installing](#installing)
* [Running](#running)
* [Swagger](#swagger)
<!--te-->

Prerequisites
============

```bash
- Linux (or WSL on windows)
- latest version of docker and docker-compose
```

Cloning repository
============

```bash
$ git clone git@github.com:raphaelsmend/clientApi.git
```
Installing
============

```bash
$ cd clientApi
$ make install
```
Running
============
for start all containers
```bash
$ make up

* this command runs the laravel server, do not close the window.
```

for tests
```bash
$ make tests
```

for restart all containers
```bash
$ make restart
```

for stop all containers
```bash
$ make down
```

Swagger
============
Swagger with all endpoints
```bash
http://localhost:8080
```