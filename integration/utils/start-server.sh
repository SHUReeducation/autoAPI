#!/usr/bin/env bash

./autoAPI -f $1 -o ./api
cd ./api || exit 1
go mod tidy && go fmt
nohup go run main.go &
sleep 5
