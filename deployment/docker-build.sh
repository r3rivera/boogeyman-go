#!/usr/bin/env bash
echo ""
echo ""
echo "Current Docker Version"
docker --version
echo ""
echo "Current Directory is $(pwd)"
echo "Current User is $USER"
echo ""
echo "Start Creating Docker Image"
chmod +x ./deployment/Dockerfile
DOCKER_TAG="boogeyman-go"

DOCKER_CTR_STOP="$(docker container stop ${DOCKER_TAG} | cut -d ':' -f 2)"
echo "Stop Status :: ${DOCKER_CTR_STOP}"

echo "Removing stopped container..."
DOCKER_CTR_RM="$(docker container rm ${DOCKER_TAG})"
echo "Remove Status :: ${DOCKER_CTR_RM}"

DOCKER_CURR="$(docker image rm ${DOCKER_TAG})"
echo "Removing existing docker images..."
echo "Remove Status :: ${DOCKER_CURR}"

GIT_COMMIT_LABEL="$(git log | head -n 1 | cut -d '/' -f 2)"

docker build -t ${DOCKER_TAG} --no-cache -f ./deployment/Dockerfile .
echo "Done Creating Docker Image"
