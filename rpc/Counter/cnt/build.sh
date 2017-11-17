#!/bin/bash


if test -e ./cnt_grpc.proto
then
    echo "Begin to build protobuf"
	protoc -I . cnt_grpc.proto --go_out=plugins=grpc:.
	echo "Success"
else
    echo 'Can not find the file [cnt_grpc.proto] !'
fi

