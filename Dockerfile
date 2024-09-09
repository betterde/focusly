ARG VERSION=latest

FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

ARG MODULE=github.com/betterde/focusly
ARG BINARY_NAME=focusly
ARG INSTALL_PATH=/usr/local/bin

RUN apk add --update gcc git make

ENV GOPATH=/tmp/buildcache
COPY . /go/src/focusly
WORKDIR /go/src/focusly
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X '${MODULE}/cmd.version=${VERSION}' -X '${MODULE}/cmd.build=`date -u`' -X '${MODULE}/cmd.commit=`git rev-parse HEAD`'" -o bin/${BINARY_NAME} main.go

FROM --platform=$BUILDPLATFORM alpine:latest

LABEL org.opencontainers.image.url="https://github.com/betterde/focusly"
LABEL org.opencontainers.image.titile="focusly"
LABEL org.opencontainers.image.vendor="Betterde Inc."
LABEL org.opencontainers.image.source="https://github.com/betterde/focusly"
LABEL org.opencontainers.image.version="${VERSION}"
LABEL org.opencontainers.image.authors="George <george@betterde.com>"
LABEL org.opencontainers.image.created="2024-08-21 20:35:00"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.description="An open source team collaboration system developed based on Go."
LABEL org.opencontainers.image.documentation="https://github.com/betterde/focusly"

COPY --from=builder /go/src/focusly/bin/focusly /usr/local/bin/focusly
RUN mkdir -p /etc/focusly
WORKDIR /root/

VOLUME ["/etc/focusly"]
ENTRYPOINT ["/usr/local/bin/focusly"]
EXPOSE 443