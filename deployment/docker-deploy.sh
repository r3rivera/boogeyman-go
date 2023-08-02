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
echo "$(cat /opt/aws/aws.txt | head -n 1)"
echo "$(cat /opt/aws/aws_key.txt | head -n 1)"
AWS_ACCESS_KEY_ID=$(cat /opt/aws/aws.txt | head -n 1)
AWS_SECRET_ACCESS_KEY=$(cat /opt/aws/aws_key.txt | head -n 1)

DOCKER_UUID="$(docker container run  \
    -p 8082:8082 \
    -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
    -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
    -d --name ${DOCKER_TAG} ${DOCKER_TAG})"
echo "New Docker ID deployed :: ${DOCKER_UUID}"