#!/bin/bash

container="gin-demo"
serverPort=80
containerPort=8080
id=$(docker ps -a | grep $container | awk '{print $1}')

if ! [ -z $id ]; then
  set -x
  docker stop $id
  docker rm $id
fi
docker run --name=$container -d -p $serverPort:$containerPort gin-demo:$1
