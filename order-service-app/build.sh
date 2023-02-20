#!/bin/bash

protoc --go_out=./proto --go_opt=paths=source_relative \
  --proto_path=./proto \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_out=require_unimplemented_servers=false:./proto \
  proto/*.proto