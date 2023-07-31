#!/usr/bin/env bash
echo ""
echo ""
echo "Current Docker Version"
docker --version
echo ""
echo "Current Directory is $(pwd)"
echo "Current User is $USER"
echo ""
echo "Stopping an existing docker container..."

DOCKER_TAG="boogeyman-go"

echo "Start deploying docker image"
echo ""
echo "Deploying a new docker container...."
DOCKER_UUID="$(docker container run -p 8082:8082 -d --name ${DOCKER_TAG} ${DOCKER_TAG})"
echo "New Docker ID deployed :: ${DOCKER_UUID}"