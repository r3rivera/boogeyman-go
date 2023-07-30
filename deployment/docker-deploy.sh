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
DOCKER_CTR_STOP="$(docker container stop ${DOCKER_TAG} | cut -d ':' -f 2)"
echo "Stop Status :: ${DOCKER_CTR_STOP}"

echo "Removing stopped container..."
DOCKER_CTR_RM="$(docker container rm ${DOCKER_TAG})"
echo "Remove Status :: ${DOCKER_CTR_RM}"

DOCKER_CURR="$(docker image rm ${DOCKER_TAG})"
echo "Removing existing docker images..."
echo "Remove Status :: ${DOCKER_CURR}"

echo "Start deploying docker image"
echo ""
echo "Deploying a new docker container...."
DOCKER_UUID="$(docker container run -p 8081:8081 -d --name ${DOCKER_TAG} ${DOCKER_TAG})"
echo "New Docker ID deployed :: ${DOCKER_UUID}"