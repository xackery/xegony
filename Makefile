default: help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
PROTO_FILES=$(shell find pb -name '*.proto')
THIS_PATH=$(shell pwd)
test: ## Run short tests
	@go test -short ./...
xegony: ## Starts xegony
	@go run main.go
proto-clean: ## Cleans up the proto directory
	@echo "pb cleanup"
	@rm -rf pb/*.pb.go
	-@rm -rf swagger/proto/*
	@mkdir -p swagger/proto
.PHONY: proto
proto: proto-clean ## Generate protobuf files
	@echo "protobuf > pb"
	@protoc \
	-I/usr/local/include \
	-I. \
	-I$(GOPATH)/src/github.com/google/protobuf/src \
	-I$(GOPATH)/src/github.com/googleapis/googleapis \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
	$(PROTO_FILES) \
	--swagger_out=logtostderr=true,allow_merge=true:swagger/ \
	--grpc-gateway_out=logtostderr=true:. \
	--go_out=plugins=grpc:.
.PHONY: swagger-ui
swagger-ui:
	@echo "visit http://localhost to see swagger documentation"
	@docker run \
	-p 80:8080 \
	-e SWAGGER_JSON=/swagger/console.json \
	-v "$(THIS_PATH)/swagger":/swagger \
	swaggerapi/swagger-ui
.PHONY: swagger-build
swagger-build: proto ## Generate swagger definitions
	@echo "swagger > swagger/xegony.swagger.json"
	@mv swagger/apidocs.swagger.json swagger/xegony.swagger.json
	@go run scripts/openapi2postman.go
.PHONY: swagger
swagger: swagger-build ## Generate swagger distribution and definitions
	@rm -rf swagger/build/*
	@docker run -v "$(shell pwd)/swagger/":/usr/src --entrypoint='' swaggerapi/swagger-codegen-cli  sh -c "java -jar /opt/swagger-codegen-cli/swagger-codegen-cli.jar generate -i /usr/src/xegony.swagger.json -l typescript-jquery -o /usr/src/build/jquery"
	@rm -rf ts/src/pb/*
	@cp -R swagger/build/jquery/* ts/src/pb/
	@rm -rf swagger/build/jquery
ts-install:
	@cd ts && npm install
ts-build:
	@cd ts && npm run build
