# syntax=docker/dockerfile:1
FROM golang:1.24.2-bookworm AS builder

WORKDIR /go/src

COPY go.mod ./
RUN go mod download
COPY ./cmd/web ./cmd/web
RUN go build -o /go/bin/web ./cmd/web/

FROM gcr.io/distroless/base-debian12:latest
COPY --from=builder /go/bin/web /
USER nonroot
ENTRYPOINT [ "/web" ]