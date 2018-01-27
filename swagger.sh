#!/bin/bash
set -e
swagger generate spec -m -o xegony.json
go run misc/openapi2postman/main.go
mv xegony.json www/docs/