#!/usr/bin/sh

VERSION="1.0.0"
APPNAME="account-srv"

# build the project
echo "Building the project..."
go build -o ${APPNAME} ./main.go

if [ ! -e ./${APPNAME} ]; then
    echo "Error: ${APPNAME} not found"
    exit 1
fi
# build image
docker build -t birdhk/${APPNAME}:${VERSION} .

# docker login
docker login --username=birdhk --password=zp521314....

# push image to docker hub
docker push birdhk/${APPNAME}:${VERSION}