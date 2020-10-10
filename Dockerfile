# Build image
FROM golang:alpine AS builder

ENV GOFLAGS="-mod=readonly"

RUN apk add --update --no-cache ca-certificates make git bash curl build-base

RUN mkdir /rituals

WORKDIR /rituals

RUN go get -u github.com/pyros2097/go-embed
RUN go get -u github.com/shiyanhui/hero/hero
RUN go get -u golang.org/x/tools/cmd/goimports

ADD ./.git     /rituals/.git
ADD ./Makefile /rituals/Makefile
ADD ./go.mod   /rituals/go.mod
ADD ./go.sum   /rituals/go.sum
ADD ./app      /rituals/app
ADD ./bin      /rituals/bin
ADD ./client   /rituals/client
ADD ./cmd      /rituals/cmd
ADD ./query    /rituals/query
ADD ./web      /rituals/web

ARG BUILD_TARGET

COPY go.* /rituals/
RUN go mod download

RUN set -xe && bash -c 'make build-release'

RUN mv build/release /build

# Final image
FROM alpine

RUN apk add --update --no-cache ca-certificates tzdata bash curl

SHELL ["/bin/bash", "-c"]

# set up nsswitch.conf for Go's "netgo" implementation
# https://github.com/gliderlabs/docker-alpine/issues/367#issuecomment-424546457
RUN test ! -e /etc/nsswitch.conf && echo 'hosts: files dns' > /etc/nsswitch.conf

ARG BUILD_TARGET

RUN if [[ "${BUILD_TARGET}" == "debug" ]]; then apk add --update --no-cache libc6-compat; fi

COPY --from=builder /build/* /rituals/

EXPOSE 6660
CMD ["/rituals/rituals"]
