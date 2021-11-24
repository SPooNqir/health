#!/bin/bash

GEN_PATH="."
GO_LIB_PATH=$(go env GOPATH)/src
GOPATH=$(go env GOPATH)

protoc \
        -I proto \
        -I $GO_LIB_PATH/include \
        --go_out=$GEN_PATH      --go_opt=paths=source_relative \
        --go-grpc_out=$GEN_PATH --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=logtostderr=true,paths=source_relative:$GEN_PATH \
        --openapiv2_out=logtostderr=true:$GEN_PATH \
        proto/health.proto

echo "mod tidy"
go mod tidy

echo "update"
go get -u ./...
