# Go Chi Clean Architecture Starter

This is a starter project for using Go Chi with Clean Architecture.

## Getting Started

To get started with the project, please create .env file based on the .env.example. and setup a few depencies like postgres and redis.

## How to Run

First, install necessary dependencies

```bash
go mod tidy

```

To run the project

```bash
go run cmd/main.go
```

To run with live reload make sure you have `nodemon` installed globally using `npm install -g nodemon`, and then run with:

```bash
make live
```

## Database Migrations

for migrations we are using golang-migrate [https://github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate) and please check `Makefile` to see available command to do Migrations.
and dont forget to change `postgresconn` value with your database url.

## To use as Project Starter

just please Find and replace `github.com/joisandresky/go-chi-clean-starter` with your project name (go modules).
