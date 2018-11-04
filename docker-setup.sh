#!/bin/bash

# build the application
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o product-service .

# build the container
docker build -t initlevel5/product-service -f Dockerfile .

# start the container
docker run -d -p 8080:8080 --name product-service initlevel5/product-service:latest
