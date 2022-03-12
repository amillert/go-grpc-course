#!/bin/bash

protoc --go_out=grpc --proto_path=grpc/proto grpc/proto/greet_msg.proto
protoc --go-grpc_out=grpc --proto_path=grpc/proto grpc/proto/greet_svc.proto
