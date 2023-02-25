#!/bin/bash

protoc --proto_path=./proto \
    --go_out=./proto --go_opt=paths=source_relative \
    --go-grpc_out=./proto --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
    --grpc-gateway_out=./proto --grpc-gateway_opt=paths=source_relative \
    proto/helloworld/helloworld.proto