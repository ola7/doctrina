#!/bin/bash

# This assumes a swarm initiatied with following network
SWARM_NET=doctrinanet

# Set env variables
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

# Build binaries
cd rsp-service;go get;go build -o rsp-service-linux-amd64;cd ..
cd healthchecker;go get; go build -o healthchecker-linux-amd64; cd ..

# Build containers
cp healthchecker/healthchecker-linux-amd64 rsp-service/
docker build -t ola7/rsp-service rsp-service/

# Clean up build artifacts
rm rsp-service/healthchecker-*

# Clean up existing services
docker service rm rsp-service
docker service rm vizualizer
docker service rm quotes-service

# Deploy services
docker service create --name=rsp-service --replicas=1 --network=$SWARM_NET -p=8989:8989 ola7/rsp-service
docker service create --name=vizualizer --replicas=1 --network=$SWARM_NET --publish=8990:8080/tcp --constraint=node.role==manager --mount=type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock dockersamples/visualizer
docker service create --name=quotes-service --replicas=1 --network=$SWARM_NET eriklupander/quotes-service
