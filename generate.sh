#!/bin/bash

protoc --go_out=grpc --proto_path=grpc/proto grpc/proto/greet_msg.proto
protoc --go-grpc_out=grpc --go-grpc_opt=require_unimplemented_servers=false --proto_path=grpc/proto grpc/proto/greet_svc.proto
