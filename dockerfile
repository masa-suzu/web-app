# Base image is in https://registry.hub.docker.com/_/golang/
# Refer to https://blog.golang.org/docker for usage
FROM golang:1.10.3-alpine3.7

ENV GOPATH $GOPATH:/go/src

RUN apk add --update git \
    && git clone https://github.com/masa-suzu/web-app.git /go/src/app \
    && cd /go/src/app \
    && go build server.go \
    && cp /go/src/app/server /go \
    && mv /go/src/app/template /go
CMD /go/server "80"