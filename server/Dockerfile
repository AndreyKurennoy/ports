ARG GOLANG_VERSION=1.11
ARG ALPINE_VERSION=3.8

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS build

RUN apk update \
    && apk add --no-cache git g++ \
    && apk add curl build-base autoconf automake libtool \
    && apk add protobuf

COPY . /build/ports/server

WORKDIR /build/ports/server

RUN export PATH=$PATH:$GOPATH/bin

RUN RUN go mod download && go mod vendor \
    && go get -u github.com/golang/protobuf/protoc-gen-go \
    && protoc --go_out=plugins=grpc:. interfaces/*.proto \
    && go build

FROM alpine:${ALPINE_VERSION}

COPY --from=build /build/ports/server/server  /opt/server-service/
COPY --from=build /build/ports/server/config.yml  /opt/server-service/

WORKDIR /opt/server-service

CMD "./server"