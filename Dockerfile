# Base image is in https://registry.hub.docker.com/_/golang/
# Refer to https://blog.golang.org/docker for usage
FROM golang:1.10.3-alpine3.7 as build

ENV GOPATH $GOPATH:/go/src

RUN apk add --update git \
    && git clone https://github.com/masa-suzu/web-app.git /go/src/app \
    && cd /go/src/app \
    && mkdir /go/bin/app \
    && go build -o /go/bin/app/server server.go \
    && ls resources \
    && cd /go/bin/app \
    && mv /go/src/app/resources .

FROM alpine:3.7
COPY --from=build /go/bin/app /
WORKDIR /
EXPOSE 80
CMD /server 80
