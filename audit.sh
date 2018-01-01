#!/bin/bash
set -e
golint api
golint bot
golint box 
golint main.go 
golint cases 
golint model 
golint storage 
golint storage/mariadb
golint web
go tool vet main.go
go tool vet api
go tool vet bot
go tool vet box
go tool vet model 
go tool vet storage 
go tool vet storage/mariadb
go tool vet web