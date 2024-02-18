LOCAL_BIN:=$(CURDIR)/bin

build-mac:
	go build -o ${LOCAL_BIN}/ ${CURDIR}/cmd/grpc_server/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${LOCAL_BIN}/service_chat_server_linux ${CURDIR}/cmd/grpc_server/main.go

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

generate:
	make generate-chat-v1-api

generate-chat-v1-api:
	mkdir -p pkg/v1/chat
	protoc --proto_path api/v1/chat \
	--go_out=pkg/v1/chat --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/v1/chat --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/v1/chat/chat.proto

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) ${LOCAL_BIN}/golangci-lint run ./... --config .golangci.pipeline.yaml