#!/bin/sh

export PORT=9000
export DATABASE_URL=postgres://postgres:example@localhost:5432/golang-starter-template
export SSL_MODE=disable

go install
go build -o bin/golang-starter-template -v .
./bin/golang-starter-template