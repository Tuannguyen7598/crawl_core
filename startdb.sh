#!/bin/bash

docker-compose up -d

sleep 5

docker exec mongo-write /scripts/rs-init.sh