#!/bin/bash

tag=$(git rev-parse --short HEAD)

set -x

go mod vendor

docker build -t gin-demo:$tag .

rm -rf vendor

./start.sh $tag

set +x