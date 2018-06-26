# Base image is in https://registry.hub.docker.com/_/golang/
# Refer to https://blog.golang.org/docker for usage
FROM golang:1.10.3-alpine3.7

ENV GOPATH $GOPATH:/go/src