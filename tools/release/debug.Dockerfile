# Content managed by Project Forge, see [projectforge.md] for details.
FROM golang:alpine

LABEL "org.opencontainers.image.authors"="Kyle U"
LABEL "org.opencontainers.image.source"="https://github.com/kyleu/rituals"
LABEL "org.opencontainers.image.vendor"="kyleu"
LABEL "org.opencontainers.image.title"="rituals.dev"
LABEL "org.opencontainers.image.description"="Work with your team to estimate work, track your progress, and gather feedback"

RUN apk add --update --no-cache ca-certificates tzdata bash curl htop libc6-compat

RUN apk add --no-cache ca-certificates dpkg gcc git musl-dev \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH"

RUN go install github.com/go-delve/delve/cmd/dlv@latest

SHELL ["/bin/bash", "-c"]

# main http port
EXPOSE 18000
# marketing port
EXPOSE 18001

WORKDIR /

ENTRYPOINT ["/rituals", "-a", "0.0.0.0"]

COPY rituals /
