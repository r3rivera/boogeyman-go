#!/usr/bin/env bash
echo "## Building the configuration of ${APP_NAME}..."
echo "## Workspace is ${WORKSPACE}"
echo "############## START ::: BUILD INFORMATION ##############"
echo ""
echo "## Displaying Golang Version..."
go version | head -n 1

echo "## Displaying GIT Version..."
git --version | head -n 1
echo ""
echo "Git URL      :: ${GIT_URL}"
echo "Git Branch   :: ${GIT_BRANCH}"
echo "Build ID     :: ${BUILD_ID}"
echo "Build Number :: ${BUILD_NUMBER}"
echo "Git Branch   :: ${GIT_BRANCH}"
GIT_COMMIT_ID="$(git log | head -n 1 | cut -d '/' -f 2)"
BUILD_DATE="$(date +%m%d%Y)"
BUILD_TIME="$(date +%H%M%S)"
echo "Build Date   :: ${BUILD_DATE}"
echo "Build Time   :: ${BUILD_TIME}"
echo "Git Commit   :: ${GIT_COMMIT_ID}"
echo ""
echo ""
echo "Start compiling the GO Application"
GO_FILE="$(find ./boogeyman | head -n 1)"
go build .
echo "Compilation Complete :: JAR File is ${GO_FILE}"
echo "############## END   ::: BUILD INFORMATION ##############"