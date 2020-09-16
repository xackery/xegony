default: help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
PROTO_FILES=$(shell find pb -name '*.proto')
THIS_PATH=$(shell pwd)
test: ## Run short tests
	@go test -short ./...
xegony: ## Starts xegony
	@go run main.go -cpuprofile cpu.prof -memprofile mem.prof
proto-clean: ## Cleans up the proto directory
	@echo "pb cleanup"
	@rm -rf pb/*.pb.go
	-@rm -rf swagger/proto/*
	@mkdir -p swagger/proto
pprof:
	go tool pprof --pdf main $(path) > file.pdf
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
	@rm -rf web/ts/src/pb/*
	@cp -R swagger/build/jquery/* web/ts/src/pb/
	@rm -rf swagger/build/jquery
web-install:
	@#ts is using stable (v10)
	@cd web/ts && npm install
	@#theme is using v8
	@cd web/theme && npm install
	@cd web/font && npm install
	@npm install -g grunt
	@gem install sass
ts-build:
	@#$(shell nvm use stable)
	@cd web/ts && npm run build
web-clean:
	rm -rf www/css/*
	rm -rf www/fonts/*
	rm -rf www/js/*
	rm -rf web/theme/dist/*
web-build: web-clean theme-build font-build ts-build
theme-build:
	@#$(shell nvm use v8.0.0)
	@cd web/theme && npm run build
	@mkdir -p www/fonts www/css www/js
	@cp web/theme/dist/js/* www/js/
	@cp web/theme/dist/css/* www/css/
	@cp web/theme/dist/fonts/* www/fonts/
font-build:
	@#$(shell nvm use v8.0.0)
	@cd web/theme && npm run build
	@cp web/font/css/* www/css/
	@cp web/font/fonts/xegonyawesome-* www/fonts/
	@rm www/fonts/*.svg
.PHONY: npm-install
npm-install:
	@(docker run --rm -v ${PWD}:/src -it xackery/webbuild:10.19.0 bash -c 'cd web/ts && npm install')
.PHONY: npm-audit-fix
npm-audit-fix:
	@(docker run --rm -v ${PWD}:/src -it xackery/webbuild:10.19.0 bash -c 'cd web/ts && npm audit fix')