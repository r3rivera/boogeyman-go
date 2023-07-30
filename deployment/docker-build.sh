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

GIT_COMMIT_LABEL="$(git log | head -n 1 | cut -d '/' -f 2)"

docker build -t ${DOCKER_TAG} --no-cache -f ./deployment/Dockerfile .
echo "Done Creating Docker Image"
