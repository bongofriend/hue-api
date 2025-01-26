FROM golang:1.23-alpine AS builder
USER root
WORKDIR /home/builder

COPY . /home/builder/hue-api
WORKDIR /home/builder/hue-api
RUN set -Eeux && \
    go mod download && \
    go mod verify

RUN go build -o app

FROM alpine:3.21.2
USER root
WORKDIR /home/app
COPY --from=builder /home/builder/hue-api/app .
ENTRYPOINT [ "./app" ]