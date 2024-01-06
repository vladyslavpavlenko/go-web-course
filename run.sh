#!/bin/zsh

go build -o bookings cmd/web/*.go
./bookings -dbname=bookings -dbuser=vladyslavpavlenko -cache=false -production=false