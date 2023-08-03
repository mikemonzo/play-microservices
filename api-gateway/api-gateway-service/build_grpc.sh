#!/bin/bash
declare -a services=("proto")
for SERVICE in "${services[@]}"; do
   DESTDIR='proto'
   mkdir -p $DESTDIR
   protoc \
       --proto_path=$SERVICE/ \
       --go_out=$DESTDIR \
       --go-grpc_out=$DESTDIR \
       $SERVICE/*.proto
done