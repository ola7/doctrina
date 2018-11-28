#!/bin/bash

SWARM_NET=doctrinanet

# RabbitMQ
docker service rm rabbitmq
docker build -t ola7/rabbitmq support/rabbitmq/
docker service create --name=rabbitmq --replicas=1 --network=$SWARM_NET -p 1883:1883 -p 5672:5672 -p 15672:15672 ola7/rabbitmq

# Spring Cloud Config
cd support/config-server
./gradlew build
cd ../..
docker build -t ola7/configserver support/config-server/
docker service rm configserver
docker service create --replicas 1 --name configserver -p 8888:8888 --network $SWARM_NET --update-delay 10s --with-registry-auth  --update-parallelism 1 ola7/configserver
