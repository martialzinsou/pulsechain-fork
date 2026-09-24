# ==============================================================================
# Dockerfile - Nœud Blockchain PulseChain Fork
# Auteur : Martial Zinsou
# ==============================================================================

FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev

COPY go.mod ./
COPY pulsechain-fork ./pulsechain-fork

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/pulsenoded ./pulsechain-fork/cmd/pulsenoded

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates curl bash

COPY --from=builder /app/pulsenoded /usr/local/bin/pulsenoded
COPY pulsechain-fork/genesis /app/genesis

EXPOSE 8545 30303

ENTRYPOINT ["/usr/local/bin/pulsenoded"]
