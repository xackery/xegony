#!/bin/bash
set -e
golint api
golint bot
golint box 
golint main.go 
golint cases 
golint cmd
golint model 
golint misc/openapi2postman
golint oauth
golint oauth/google
golint storage 
golint storage/file
golint storage/mariadb
golint storage/memory
golint web
go tool vet api
go tool vet bot
go tool vet box 
go tool vet main.go 
go tool vet cases 
go tool vet cmd
go tool vet model 
go tool vet misc/openapi2postman
go tool vet oauth
go tool vet oauth/google
go tool vet storage 
go tool vet storage/file
go tool vet storage/mariadb
go tool vet storage/memory
go tool vet web