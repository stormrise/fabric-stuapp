#!/bin/bash
# Exit on first error
set -e

cd ../basic-network
docker-compose -f docker-compose.yml stop

docker-compose -f docker-compose.yml kill && docker-compose -f docker-compose.yml down

rm -f ~/.hfc-key-store/*

echo y | docker network prune

docker rmi -f $(docker images dev-* -q)

echo =====================================================
echo ===================clean success!====================
echo =====================================================
