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
echo "AWS $(aws --profile default configure get aws_access_key_id)"
AWS_ACCESS_KEY_ID=$(aws --profile default configure get aws_access_key_id)
AWS_SECRET_ACCESS_KEY=$(aws --profile default configure get aws_secret_access_key)

DOCKER_UUID="$(docker container run  \
    -p 8082:8082 \
    -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
    -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
    -d --name ${DOCKER_TAG} ${DOCKER_TAG})"
echo "New Docker ID deployed :: ${DOCKER_UUID}"